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
	"math/big"

	"github.com/Dacchain/go-Dacchain/common"
	"github.com/Dacchain/go-Dacchain/consensus"
	"github.com/Dacchain/go-Dacchain/core/types"
	"github.com/Dacchain/go-Dacchain/core/vm"
	"github.com/Dacchain/go-Dacchain/params"
	"github.com/Dacchain/go-Dacchain/log"
)

// ChainContext supports retrieving headers and consensus parameters from the
// current blockchain to be used during transaction processing.
type ChainContext interface {
	// Engine retrieves the chain's consensus engine.
	Engine() consensus.Engine

	// GetHeader returns the hash corresponding to their hash.
	GetHeader(common.Hash, uint64) *types.Header

	GetDelegatePoll() (*map[common.Address]types.Candidate, error)

	GetGenesisConfig() *params.ChainConfig
}

// NewEVMContext creates a new context for use in the EVM.
func NewEVMContext(msg Message, header *types.Header, chain ChainContext, author *common.Address) vm.Context {
	var delegates *map[common.Address]types.Candidate
	var err error
	if msg.Action() == types.ActionRegister || msg.Action() == types.ActionAddVote || msg.Action() == types.ActionSubVote {
		delegates, err = chain.GetDelegatePoll()
		if err != nil {
			return vm.Context{}
		}
	}

	return vm.Context{
		CanTransfer:  CanTransfer,
		Transfer:     Transfer,
		Vote:         Vote,
		GetHash:      GetHashFn(header, chain),
		Origin:       msg.From(),
		Coinbase:     header.Coinbase,
		BlockNumber:  new(big.Int).Set(header.Number),
		Time:         new(big.Int).Set(header.Time),
		Difficulty:   new(big.Int).Set(types.BlockDifficult),
		GasLimit:     header.GasLimit,
		GasPrice:     new(big.Int).Set(msg.GasPrice()),
		DelegateList: delegates,
	}
}

// GetHashFn returns a GetHashFunc which retrieves header hashes by number
func GetHashFn(ref *types.Header, chain ChainContext) func(n uint64) common.Hash {
	var cache map[uint64]common.Hash

	return func(n uint64) common.Hash {
		// If there's no hash cache yet, make one
		if cache == nil {
			cache = map[uint64]common.Hash{
				ref.Number.Uint64() - 1: ref.ParentHash,
			}
		}
		// Try to fulfill the request from the cache
		if hash, ok := cache[n]; ok {
			return hash
		}
		// Not cached, iterate the blocks and cache the hashes
		for header := chain.GetHeader(ref.ParentHash, ref.Number.Uint64()-1); header != nil; header = chain.GetHeader(header.ParentHash, header.Number.Uint64()-1) {
			cache[header.Number.Uint64()-1] = header.ParentHash
			if n == header.Number.Uint64()-1 {
				return header.ParentHash
			}
		}
		return common.Hash{}
	}
}

// CanTransfer checks wether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db vm.StateDB, addr common.Address, asset *common.Address, amount *big.Int) bool {
	if asset == nil {
		return db.GetBalance(addr).Cmp(amount) >= 0
	} else {
		return db.GetAssetBalance(addr, *asset).Cmp(amount) >= 0
	}
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db vm.StateDB, sender, recipient common.Address, asset *common.Address, amount *big.Int) {
	if asset == nil {
		db.SubBalance(sender, amount)
		db.AddBalance(recipient, amount)
	} else {
		ok := db.SubAssetBalance(sender, *asset, amount)
		if ok {
			db.AddAssetBalance(recipient, *asset, amount)
		}
	}
}

func Vote(db vm.StateDB, user common.Address, amount *big.Int, vote []types.Vote, delegateList *map[common.Address]types.Candidate, maxElectDelegate int64) error {
	d := *delegateList
	//for _, v := range vote {
	//	if _, ok := d[strings.ToLower(v.Candidate.Hex())]; !ok {
	//		return ErrVoteList
	//	}
	//}
	newVoteList, diff, err := changeVoteList(db.GetVoteList(user), vote, d)
	if err != nil {
		log.Debug("InVoteError", "err", err)
		return err
	}
	lockBalance, _ := new(big.Int).SetString(db.GetLockBalance(user).String(), 0)
	//log.Info("Info", "lockBalance", lockBalance)
	if (lockBalance.Add(lockBalance, diff)).Cmp(new(big.Int).Mul(big.NewInt(params.Dac), big.NewInt(maxElectDelegate))) > 0 {
		return vm.ErrVote
	}
	//log.Info("CoreEVMN", "diff", diff,"balance", db.GetBalance(user), "lockBalance", db.GetLockBalance(user))
	db.SubBalance(user, diff)
	//log.Info("LockBalance","balance", db.GetBalance(user), "lockBalance", db.GetLockBalance(user))
	db.AddLockBalance(user, diff)
	//log.Info("SetVOteList", "voteList", newVoteList)
	db.SetVoteList(user, newVoteList)
	//log.Info("Setvote finish", "balance", db.GetBalance(user), "lockBalance", db.GetLockBalance(user))
	return nil
}

func changeVoteList(prevVoteList []common.Address, curVoteList []types.Vote, delegateList map[common.Address]types.Candidate) ([]common.Address, *big.Int, error) {
	var (
		voteChangeList = prevVoteList
		diff           = int64(0)
	)
	//log.Info("core|evm|Vote", "dbVoteList", prevVoteList,"newVoteList", curVoteList)
	for _, vote := range curVoteList {
		switch vote.Operation {
		case 0:
			if _, contain := sliceContains(*vote.Candidate, voteChangeList); !contain {
				if _, ok := delegateList[*vote.Candidate]; !ok {
					return nil, nil, vm.ErrVote
				}
				diff += 1
				voteChangeList = append(voteChangeList, *vote.Candidate)
			} else {
				return nil, big.NewInt(0), vm.ErrVote
			}
		case 1:
			if j, contain := sliceContains(*vote.Candidate, voteChangeList); contain {
				diff -= 1
				voteChangeList = append(voteChangeList[:j], voteChangeList[j+1:]...)
			} else {
				return nil, big.NewInt(0), vm.ErrVote
			}
		default:
			return nil, big.NewInt(0), vm.ErrVote
		}
	}
	lockNumber := big.NewInt(diff)
	lockNumber = lockNumber.Mul(lockNumber, big.NewInt(params.Dac))
	return voteChangeList, lockNumber, nil
}

func sliceContains(address common.Address, slice []common.Address) (int, bool) {
	for i, a := range slice {
		//log.Info("address == a", "address", address.Hex(), "a", a.Hex(), "result", address == a)
		if address == a {
			return i, true
		}
	}
	return -1, false
}
