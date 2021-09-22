package fccCoin

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
)

type client struct {
	PublicKey  crypto.PublicKey
	privateKey crypto.PrivateKey
}

func NewClient() (*client, error) {
	var privateKey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("could not generate Private Key for RSA")
	}
	var publicKey = privateKey.Public()
	return &client{
		publicKey,
		privateKey,
	}, nil
}
