// Copyright 2018 The go-Dacchain Authors
// This file is part of the go-Dacchain library.
//
// The go-Dacchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-Dacchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-Dacchain library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Dacchain/go-Dacchain/dacdb"
	"github.com/Dacchain/go-Dacchain/common"
	"github.com/Dacchain/go-Dacchain/common/hexutil"
	"github.com/Dacchain/go-Dacchain/common/math"
	"github.com/Dacchain/go-Dacchain/consensus/delegatestate"
	"github.com/Dacchain/go-Dacchain/core/state"
	"github.com/Dacchain/go-Dacchain/core/types"
	"github.com/Dacchain/go-Dacchain/crypto/sha3"
	"github.com/Dacchain/go-Dacchain/log"
	"github.com/Dacchain/go-Dacchain/params"
	"github.com/Dacchain/go-Dacchain/rlp"
	"github.com/Dacchain/go-Dacchain/util"
	"math/big"
	"strings"
)

//go:generate gencodec -type Genesis -field-override genesisSpecMarshaling -out gen_genesis.go
//go:generate gencodec -type GenesisAccount -field-override genesisAccountMarshaling -out gen_genesis_account.go
//go:generate go run genMkAgentsFile.go -number 6 -password yujian -keystore-dir .

var errGenesisNoConfig = errors.New("genesis has no chain configuration")

const genesisExtra = "DAC genesis"

// Genesis specifies the header fields, state of a genesis block. It also defines hard
// fork switch-over blocks through the chain configuration.
type Genesis struct {
	Config    *params.ChainConfig `json:"config"`
	Timestamp uint64              `json:"timestamp"`
	ExtraData []byte              `json:"extraData"`
	GasLimit  uint64              `json:"gasLimit"   gencodec:"required"`
	Coinbase  common.Address      `json:"coinbase"`
	Alloc     GenesisAlloc        `json:"alloc"      gencodec:"required"`
	Agents    GenesisAgents       `json:"agents"     gencodec:"required"`
	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     uint64      `json:"number"`
	GasUsed    uint64      `json:"gasUsed"`
	ParentHash common.Hash `json:"parentHash"`
}

// GenesisAlloc specifies the initial state that is part of the genesis block.
type GenesisAlloc map[common.Address]GenesisAccount

type GenesisAgents []types.Candidate

func (ga *GenesisAlloc) UnmarshalJSON(data []byte) error {
	m := make(map[common.UnprefixedAddress]GenesisAccount)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*ga = make(GenesisAlloc)
	for addr, a := range m {
		(*ga)[common.Address(addr)] = a
	}
	return nil
}

// GenesisAccount is an account in the state of the genesis block.
type GenesisAccount struct {
	Code       []byte                      `json:"code,omitempty"`
	Storage    map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance    *big.Int                    `json:"balance" gencodec:"required"`
	Nonce      uint64                      `json:"nonce,omitempty"`
	PrivateKey []byte                      `json:"secretKey,omitempty"` // for tests
}

// field type overrides for gencodec
type genesisSpecMarshaling struct {
	Timestamp math.HexOrDecimal64
	ExtraData hexutil.Bytes
	GasLimit  math.HexOrDecimal64
	GasUsed   math.HexOrDecimal64
	Number    math.HexOrDecimal64
	Alloc     map[common.UnprefixedAddress]GenesisAccount
}

type genesisAccountMarshaling struct {
	Code       hexutil.Bytes
	Balance    *math.HexOrDecimal256
	Nonce      math.HexOrDecimal64
	Storage    map[storageJSON]storageJSON
	PrivateKey hexutil.Bytes
}

// storageJSON represents a 256 bit byte array, but allows less than 256 bits when
// unmarshaling from hex.
type storageJSON common.Hash

func (h *storageJSON) UnmarshalText(text []byte) error {
	text = bytes.TrimPrefix(text, []byte("0x"))
	if len(text) > 64 {
		return fmt.Errorf("too many hex characters in storage key/value %q", text)
	}
	offset := len(h) - len(text)/2 // pad on the left
	if _, err := hex.Decode(h[offset:], text); err != nil {
		fmt.Println(err)
		return fmt.Errorf("invalid hex storage key/value %q", text)
	}
	return nil
}

func (h storageJSON) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

// GenesisMismatchError is raised when trying to overwrite an existing
// genesis block with an incompatible one.
type GenesisMismatchError struct {
	Stored, New common.Hash
}

func (e *GenesisMismatchError) Error() string {
	return fmt.Sprintf("database already contains an incompatible genesis block (have %x, new %x)", e.Stored[:8], e.New[:8])
}

// SetupGenesisBlock writes or updates the genesis block in db.
// The block that will be used is:
//
//                          genesis == nil       genesis != nil
//                       +------------------------------------------
//     db has no genesis |  main-net default  |  genesis
//     db has genesis    |  from DB           |  genesis (if compatible)
//
// The stored chain configuration will be updated if it is compatible (i.e. does not
// specify a fork block below the local head block). In case of a conflict, the
// error is a *params.ConfigCompatError and the new, unwritten config is returned.
//
// The returned chain configuration is never nil.
func SetupGenesisBlock(db dacdb.Database, genesis *Genesis) (*params.ChainConfig, common.Hash, *Genesis, error) {
	if genesis != nil && genesis.Config == nil {
		return params.AllDavinciProtocolChanges, common.Hash{}, genesis, errGenesisNoConfig
	}

	// Just commit the new block if there is no stored genesis block.
	stored := GetCanonicalHash(db, 0)
	if (stored == common.Hash{}) {
		if genesis == nil {
			log.Info("Writing default main-net genesis block")
			genesis = DefaultGenesisBlock()
		} else {
			log.Info("Writing custom genesis block")
		}
		block, err := genesis.Commit(db)
		return genesis.Config, block.Hash(), genesis, err
	}

	// Check whether the genesis block is already written.
	if genesis != nil {
		block, _, _ := genesis.ToBlock()
		hash := block.Hash()
		if hash != stored {
			return genesis.Config, block.Hash(), genesis, &GenesisMismatchError{stored, hash}
		}
	}

	// Get the existing chain configuration.
	newcfg := genesis.configOrDefault(stored)
	storedcfg, err := GetChainConfig(db, stored)
	if err != nil {
		if err == ErrChainConfigNotFound {
			// This case happens if a genesis write was interrupted.
			log.Warn("Found genesis block without chain config")
			err = WriteChainConfig(db, stored, newcfg)
		}
		return newcfg, stored, genesis, err
	}
	// Special case: don't change the existing config of a non-mainnet chain if no new
	// config is supplied. These chains would get AllProtocolChanges (and a compat error)
	// if we just continued here.
	if genesis == nil && stored != params.MainnetGenesisHash {
		return storedcfg, stored, genesis, nil
	}
	// Check config compatibility and write the config. Compatibility errors
	// are returned to the caller unless we're already at block zero.
	height := GetBlockNumber(db, GetHeadHeaderHash(db))
	if height == missingNumber {
		return newcfg, stored, genesis, fmt.Errorf("missing block number for head header hash")
	}
	compatErr := storedcfg.CheckCompatible(newcfg, height)
	if compatErr != nil && height != 0 && compatErr.RewindTo != 0 {
		return newcfg, stored, genesis, compatErr
	}
	return newcfg, stored, genesis, WriteChainConfig(db, stored, newcfg)
}

func (g *Genesis) configOrDefault(ghash common.Hash) *params.ChainConfig {
	switch {
	case g != nil:
		return g.Config
	case ghash == params.MainnetGenesisHash:
		return params.MainnetChainConfig
	case ghash == params.TestnetGenesisHash:
		return params.TestnetChainConfig
	default:
		return params.AllDavinciProtocolChanges
	}
}

// ToBlock creates the block and state of a genesis specification.
func (g *Genesis) ToBlock() (*types.Block, *state.StateDB, *delegatestate.DelegateDB) {
	db, _ := dacdb.NewMemDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db))
	delegatedb, _ := delegatestate.New(common.Hash{}, delegatestate.NewDatabase(db))
	for addr, account := range g.Alloc {
		statedb.AddBalance(addr, account.Balance)
		statedb.SetCode(addr, account.Code)
		statedb.SetNonce(addr, account.Nonce)
		for key, value := range account.Storage {
			statedb.SetState(addr, key, value)
		}
	}
	root := statedb.IntermediateRoot(false)
	// add agents
	for _, agent := range g.Agents {
		addr := common.HexToAddress(agent.Address)
		obj := delegatedb.GetOrNewStateObject(addr, agent.Nickname, agent.RegisterTime)
		obj.AddVote(big.NewInt(int64(agent.Vote)))
	}
	delegateRoot := delegatedb.IntermediateRoot(false)
	topDelegates := delegatedb.GetDelegates()
	config := g.Config
	MaxElectDelegate := config.MaxElectDelegate.Int64()
	if len(topDelegates) > int(MaxElectDelegate) {
		topDelegates = topDelegates[:int(MaxElectDelegate)]
	}
	shuffleNewRound := util.ShuffleNewRound(int64(g.Timestamp), int(MaxElectDelegate), topDelegates, config.BlockInterval.Int64())
	shuffleList := types.ShuffleList{ShuffleDels: shuffleNewRound}
	rlpShufflehash := rlpHash(shuffleList)

	encodeBytes := hexutil.Encode([]byte("DAC official"))
	agentName := hexutil.MustDecode(encodeBytes)

	head := &types.Header{
		Number:             new(big.Int).SetUint64(g.Number),
		Time:               new(big.Int).SetUint64(g.Timestamp),
		ParentHash:         g.ParentHash,
		Extra:              g.ExtraData,
		GasLimit:           g.GasLimit,
		GasUsed:            g.GasUsed,
		Coinbase:           g.Coinbase,
		Root:               root,
		DelegateRoot:       delegateRoot,
		ShuffleHash:        rlpShufflehash,
		AgentName:          agentName,
		ShuffleBlockNumber: big.NewInt(int64(g.Number)),
	}
	if g.GasLimit == 0 {
		head.GasLimit = params.GenesisGasLimit
	}
	return types.NewBlock(head, nil, nil), statedb, delegatedb
}

// Commit writes the block and state of a genesis specification to the database.
// The block is committed as the canonical head block.
func (g *Genesis) Commit(db dacdb.Database) (*types.Block, error) {
	block, statedb, delegatedb := g.ToBlock()
	if block.Number().Sign() != 0 {
		return nil, fmt.Errorf("can't commit genesis block with number > 0")
	}
	if _, err := statedb.CommitTo(db, false); err != nil {
		return nil, fmt.Errorf("cannot write state: %v", err)
	}
	if _, err := delegatedb.CommitTo(db, false); err != nil {
		return nil, fmt.Errorf("cannot write delegate state: %v", err)
	}
	if err := WriteTd(db, block.Hash(), block.NumberU64(), types.BlockDifficult); err != nil {
		return nil, err
	}
	if err := WriteBlock(db, block); err != nil {
		return nil, err
	}
	if err := WriteBlockReceipts(db, block.Hash(), block.NumberU64(), nil); err != nil {
		return nil, err
	}
	if err := WriteCanonicalHash(db, block.Hash(), block.NumberU64()); err != nil {
		return nil, err
	}
	if err := WriteHeadBlockHash(db, block.Hash()); err != nil {
		return nil, err
	}
	if err := WriteHeadHeaderHash(db, block.Hash()); err != nil {
		return nil, err
	}
	config := g.Config
	if config == nil {
		config = params.AllDavinciProtocolChanges
	}
	return block, WriteChainConfig(db, block.Hash(), config)
}

// MustCommit writes the genesis block and state to db, panicking on error.
// The block is committed as the canonical head block.
func (g *Genesis) MustCommit(db dacdb.Database) *types.Block {
	block, err := g.Commit(db)
	if err != nil {
		panic(err)
	}
	return block
}

// GenesisBlockForTesting creates and writes a block in which addr has the given wei balance.
func GenesisBlockForTesting(db dacdb.Database, addr common.Address, balance *big.Int) *types.Block {
	g := Genesis{Alloc: GenesisAlloc{addr: {Balance: balance}}}
	return g.MustCommit(db)
}

// DefaultGenesisBlock returns the Davinci main net genesis block.
func DefaultGenesisBlock() *Genesis {
	encodeBytes := hexutil.Encode([]byte(genesisExtra))
	agentName := hexutil.MustDecode(encodeBytes)
	return &Genesis{
		Config:    params.MainnetChainConfig,
		ExtraData: agentName,
		GasLimit:  params.GenesisGasLimit,
		Alloc:     decodePrealloc(mainnetAllocData),
		Agents:    decodeGenesisAgents(mainnetAgentData),
		Timestamp: 0,
	}
}

// DefaultTestnetGenesisBlock returns the Ropsten network genesis block.
func DefaultTestnetGenesisBlock() *Genesis {
	encodeBytes := hexutil.Encode([]byte(genesisExtra))
	agentName := hexutil.MustDecode(encodeBytes)
	return &Genesis{
		Config:    params.TestnetChainConfig,
		ExtraData: agentName,
		GasLimit:  250000000,
		Alloc:     decodePrealloc(testnetAllocData),
		Agents:    decodeGenesisAgents(testnetAgentData),
		Timestamp: 0,
	}
}

// DefaultRinkebyGenesisBlock returns the Rinkeby network genesis block.
func DefaultRinkebyGenesisBlock() *Genesis {
	encodeBytes := hexutil.Encode([]byte(genesisExtra))
	agentName := hexutil.MustDecode(encodeBytes)
	return &Genesis{
		Config:    params.AllDavinciProtocolChanges,
		Timestamp: 1492009146,
		ExtraData: agentName,
		GasLimit:  4700000,
		Alloc:     decodePrealloc(rinkebyAllocData),
	}
}

// DeveloperGenesisBlock returns the 'dac --dev' genesis block. Note, this must
// be seeded with the
func DeveloperGenesisBlock(developer common.Address) *Genesis {
	// Override the default period to the user requested one
	config := *params.AllDavinciProtocolChanges
	encodeBytes := hexutil.Encode([]byte(genesisExtra))
	agentName := hexutil.MustDecode(encodeBytes)
	var geneAgents GenesisAgents
	candidates := make([]types.Candidate, 0)
	candidates = append(candidates, types.Candidate{Address: strings.ToLower(developer.Hex()), Vote: 1, Nickname: developer.Hex(), RegisterTime: uint64(11111112111)})
	geneAgents = append(geneAgents, candidates...)
	// Assemble and return the genesis with the precompiles and faucet pre-funded
	return &Genesis{
		Config:    &config,
		ExtraData: agentName,
		GasLimit:  250000000,
		Alloc: map[common.Address]GenesisAccount{
			developer: {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
		},
		Agents:    geneAgents,
		Timestamp: 1492009146,
	}
}

func decodePrealloc(data string) GenesisAlloc {
	var p []struct{ Addr, Balance *big.Int }
	if err := rlp.NewStream(strings.NewReader(data), 0).Decode(&p); err != nil {
		panic(err)
	}
	ga := make(GenesisAlloc, len(p))
	for _, account := range p {
		ga[common.BigToAddress(account.Addr)] = GenesisAccount{Balance: account.Balance}
	}
	return ga
}

func decodeGenesisAgents(data string) GenesisAgents {
	var p GenesisAgents
	if err := rlp.NewStream(strings.NewReader(data), 0).Decode(&p); err != nil {
		panic(err)
	}
	return p
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
