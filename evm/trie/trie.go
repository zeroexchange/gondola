// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package trie

import (

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type Trie struct {
	t         *trie.Trie
	trieRoots []common.Hash
}

func NewTrie(root common.Hash, db *Database) (*Trie, error) {

	// TODO: look into cache values
	// this creates a new trie database with our KVDB as the diskDB for node storage
	newTrie, err := trie.New(root, trie.NewDatabaseWithCache(db, 0))

	trie := &Trie{

		t: newTrie,
	}

	if root != (common.Hash{}) && root != trie.emptyRoot {
		trie.trieRoots = []common.Hash{root}
	}

	return trie
}

func (*Trie) UpdateTrie(transactions []common.Hash, transactionRoot common.Hash) {
	for i, tx := range transactions {
		key = rlp.EncodeToBytes(i)
		Trie.t.Update(key, tx)
	}

	require(t.root == transactionRoot, "transaction roots don't match")

	Trie.trieRoots = append(Trie.trieRoots, transactionRoot)
}

// retrieves a proof from a trie object, given the root of the trie it is contained in and the key
func (t *Trie) RetrieveProof(root common.Hash, key []byte) *ProofDatabase {
	var proof = *ProofDatabase
	err := t.t.Prove(key, 0, &proof)
	if err != nil {
		return proof
	}
}
