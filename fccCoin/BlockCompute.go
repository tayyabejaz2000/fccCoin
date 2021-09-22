package fccCoin

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const difficulty = 4

func VerifyingProof(last_proof int64, proof_no int64) bool {
	var blob = bytes.NewBufferString(fmt.Sprint(last_proof, proof_no))
	var guess_hash = sha256.Sum256(blob.Bytes())
	return hex.EncodeToString(guess_hash[:])[:4] == strings.Repeat("0", difficulty)
}

func ProofOfWork(last_proof int64) int64 {
	var proof_no int64 = 0
	for !VerifyingProof(proof_no, last_proof) {
		proof_no++
	}
	return proof_no
}
