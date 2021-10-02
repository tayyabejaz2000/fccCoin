package fccCoin

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const difficulty = 4

func VerifyingProof(last_hash [32]byte, proof_no int64) bool {
	var blob = fmt.Sprint(last_hash[:], proof_no)
	var guess_hash = sha256.Sum256([]byte(blob))
	return hex.EncodeToString(guess_hash[:])[:4] == strings.Repeat("0", difficulty)
}

func ProofOfWork(last_hash [32]byte) int64 {
	var proof_no int64 = 0
	for !VerifyingProof(last_hash, proof_no) {
		proof_no++
	}
	return proof_no
}
