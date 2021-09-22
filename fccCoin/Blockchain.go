package fccCoin

import (
	"container/list"
	"time"
)

type blockchain struct {
	blocks       *list.List
	current_data []*transaction
}

func CreateBlockchain() *blockchain {
	var b = blockchain{list.New(), nil}
	b.CreateGenesisBlock()
	return &b
}

func (b *blockchain) AddBlock(proof_no int64, prev_hash [32]byte) *block {
	var blk = CreateBlock(
		b.blocks.Len(),
		proof_no,
		prev_hash,
		b.current_data,
		time.Now(),
	)
	b.current_data = nil
	b.blocks.PushBack(blk)
	return blk
}

func (b *blockchain) CreateGenesisBlock() {
	b.AddBlock(0, [32]byte{})
}

func (b *blockchain) NewData(sender string, reciever string, quantity int64) {
	b.current_data = append(b.current_data, CreateTransaction(sender, reciever, quantity))
}

func (b *blockchain) LatestBlock() *block {
	var it = b.blocks.Back()
	if it == nil {
		return nil
	}
	return it.Value.(*block)
}

func (b *blockchain) BlockMining(miner_details string) *block {
	b.NewData("0", miner_details, 1)
	var lastBlock = b.LatestBlock()

	var lastProofNo = lastBlock.proof_no
	var proof_no = ProofOfWork(lastProofNo)

	var lastHash = lastBlock.BlockHash()
	var block = b.AddBlock(proof_no, lastHash)

	return block
}
