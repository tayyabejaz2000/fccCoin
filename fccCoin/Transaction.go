package fccCoin

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type transaction struct {
	Sender   string `json:"sender,omitempty"`
	Reciever string `json:"reciever,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
}

func CreateTransaction(sender string, reciever string, quantity int64) *transaction {
	return &transaction{
		Sender:   sender,
		Reciever: reciever,
		Quantity: quantity,
	}
}

func (t *transaction) TransactionHash() [32]byte {
	var blob = bytes.NewBufferString(fmt.Sprint(t.Sender, t.Reciever, t.Quantity))
	var hash = sha256.Sum256(blob.Bytes())
	return hash
}
