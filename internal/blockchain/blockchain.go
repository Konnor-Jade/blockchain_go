package blockchain

import (
	"time"
)

func GenerateBlock(prevBlock Block, transactions []Transaction, difficulty int) Block {
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
	}
	// 工作量证明
	pow := NewProofOfWork(&newBlock, difficulty)
	hash, nonce := pow.Run()
	newBlock.Hash = hash
	newBlock.Nonce = nonce
	return newBlock
}
