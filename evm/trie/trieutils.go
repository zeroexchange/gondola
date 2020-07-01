// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

import (

	"fmt"

	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

func VerifyProof(root common.Hash, key []bytes, proof *ProofDatabase) bool {
	exists, err := trie.VerifyProof(root, key, proof)

	if err != nil {
		return false, err
	}

	return true, nil
}