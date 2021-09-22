package main

import (
	"fmt"
	"main/fccCoin"
)

func main() {
	var blockchain = fccCoin.CreateBlockchain()
	blockchain.BlockMining("Tobii")

	fmt.Println("Block Added: ", blockchain.LatestBlock().String())
}
