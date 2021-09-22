package fccCoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type block struct {
	index     int
	proof_no  int64
	prev_hash [32]byte
	data      []transaction
	timestamp time.Time
}

func CreateBlock(index int, proof_no int64, prev_hash [32]byte, data []*transaction, timestamp time.Time) *block {
	var b = block{
		index:     index,
		proof_no:  proof_no,
		prev_hash: prev_hash,
		data:      make([]transaction, len(data)),
		timestamp: timestamp,
	}
	for i, t := range data {
		b.data[i] = *t
	}
	return &b
}

func (b *block) SetData(data []*transaction) {
	b.data = make([]transaction, len(data))
	for i, t := range data {
		b.data[i] = *t
	}
}

func (b *block) BlockHash() [32]byte {
	var blob = bytes.NewBufferString(fmt.Sprint(b.index, b.proof_no, b.prev_hash, b.data, b.timestamp))
	var hash = sha256.Sum256(blob.Bytes())
	return hash
}

func (b *block) GetProofNo() int64 {
	return b.proof_no
}

func (b *block) String() string {
	return fmt.Sprintf("Index: %d\nProof No: %d\nPrevious Hash: %s\nData: %v\nTimestamp: %s\n",
		b.index,
		b.proof_no,
		strings.ToUpper(hex.EncodeToString(b.prev_hash[:])),
		b.data,
		b.timestamp.String(),
	)
}

func (curr_block *block) CheckValidity(prev_block *block) bool {
	if prev_block.index+1 != curr_block.index {
		return false
	} else if prev_block.BlockHash() != curr_block.prev_hash {
		return false
	} else if !VerifyingProof(curr_block.proof_no, prev_block.proof_no) {
		return false
	} else if curr_block.timestamp.Before(prev_block.timestamp) {
		return false
	}
	return true
}
