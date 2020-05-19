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

// Package light implements on-demand retrieval capable state and chain objects
// for the Davinci Light Client.
package light

import (
	"context"
	"math/big"

	"github.com/Dacchain/go-Dacchain/common"
	"github.com/Dacchain/go-Dacchain/core"
	"github.com/Dacchain/go-Dacchain/core/types"
	"github.com/Dacchain/go-Dacchain/dacdb"

)

// NoOdr is the default context passed to an ODR capable function when the ODR
// service is not required.
var NoOdr = context.Background()

// OdrBackend is an interface to a backend service that handles ODR retrievals type
type OdrBackend interface {
	Database() dacdb.Database
	ChtIndexer() *core.ChainIndexer
	BloomTrieIndexer() *core.ChainIndexer
	BloomIndexer() *core.ChainIndexer
	Retrieve(ctx context.Context, req OdrRequest) error
}

// OdrRequest is an interface for retrieval requests
type OdrRequest interface {
	StoreResult(db dacdb.Database)
}

// TrieID identifies a state or account storage trie
type TrieID struct {
	BlockHash, Root common.Hash
	BlockNumber     uint64
	AccKey          []byte
}

// StateTrieID returns a TrieID for a state trie belonging to a certain block
// header.
func StateTrieID(header *types.Header) *TrieID {
	return &TrieID{
		BlockHash:   header.Hash(),
		BlockNumber: header.Number.Uint64(),
		AccKey:      nil,
		Root:        header.Root,
	}
}

// StorageTrieID returns a TrieID for a contract storage trie at a given account
// of a given state trie. It also requires the root hash of the trie for
// checking Merkle proofs.
func StorageTrieID(state *TrieID, addrHash, root common.Hash) *TrieID {
	return &TrieID{
		BlockHash:   state.BlockHash,
		BlockNumber: state.BlockNumber,
		AccKey:      addrHash[:],
		Root:        root,
	}
}

// TrieRequest is the ODR request type for state/storage trie entries
type TrieRequest struct {
	OdrRequest
	Id    *TrieID
	Key   []byte
	Proof *NodeSet
}

// StoreResult stores the retrieved data in local database
func (req *TrieRequest) StoreResult(db dacdb.Database) {
	req.Proof.Store(db)
}

// CodeRequest is the ODR request type for retrieving contract code
type CodeRequest struct {
	OdrRequest
	Id   *TrieID // references storage trie of the account
	Hash common.Hash
	Data []byte
}

// StoreResult stores the retrieved data in local database
func (req *CodeRequest) StoreResult(db dacdb.Database) {
	db.Put(req.Hash[:], req.Data)
}

// BlockRequest is the ODR request type for retrieving block bodies
type BlockRequest struct {
	OdrRequest
	Hash   common.Hash
	Number uint64
	Rlp    []byte
}

// StoreResult stores the retrieved data in local database
func (req *BlockRequest) StoreResult(db dacdb.Database) {
	core.WriteBodyRLP(db, req.Hash, req.Number, req.Rlp)
}

// ReceiptsRequest is the ODR request type for retrieving block bodies
type ReceiptsRequest struct {
	OdrRequest
	Hash     common.Hash
	Number   uint64
	Receipts types.Receipts
}

// StoreResult stores the retrieved data in local database
func (req *ReceiptsRequest) StoreResult(db dacdb.Database) {
	core.WriteBlockReceipts(db, req.Hash, req.Number, req.Receipts)
}

// ChtRequest is the ODR request type for state/storage trie entries
type ChtRequest struct {
	OdrRequest
	ChtNum, BlockNum uint64
	ChtRoot          common.Hash
	Header           *types.Header
	Td               *big.Int
	Proof            *NodeSet
}

// StoreResult stores the retrieved data in local database
func (req *ChtRequest) StoreResult(db dacdb.Database) {
	// if there is a canonical hash, there is a header too
	core.WriteHeader(db, req.Header)
	hash, num := req.Header.Hash(), req.Header.Number.Uint64()
	core.WriteTd(db, hash, num, req.Td)
	core.WriteCanonicalHash(db, hash, num)
}

// BloomRequest is the ODR request type for retrieving bloom filters from a CHT structure
type BloomRequest struct {
	OdrRequest
	BloomTrieNum   uint64
	BitIdx         uint
	SectionIdxList []uint64
	BloomTrieRoot  common.Hash
	BloomBits      [][]byte
	Proofs         *NodeSet
}

// StoreResult stores the retrieved data in local database
func (req *BloomRequest) StoreResult(db dacdb.Database) {
	for i, sectionIdx := range req.SectionIdxList {
		sectionHead := core.GetCanonicalHash(db, (sectionIdx+1)*BloomTrieFrequency-1)
		// if we don't have the canonical hash stored for this section head number, we'll still store it under
		// a key with a zero sectionHead. GetBloomBits will look there too if we still don't have the canonical
		// hash. In the unlikely case we've retrieved the section head hash since then, we'll just retrieve the
		// bit vector again from the network.
		core.WriteBloomBits(db, req.BitIdx, sectionIdx, sectionHead, req.BloomBits[i])
	}
}
