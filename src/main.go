package main

import (
	"blockchain_go/src/blockchain"
	"log/slog"
	"time"
)

func main() {
	genesisBlock := blockchain.Block{
		Index:     0,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "创世区块",
		PrevHash:  "",
		Hash:      "",
		Nonce:     0,
	}
	genesisBlock.Hash, genesisBlock.Nonce = blockchain.NewProofOfWork(&genesisBlock, 6).Run()
	blockChain := []blockchain.Block{genesisBlock}
	blockChain = append(blockChain, blockchain.GenerateBlock(
		blockChain[len(blockChain)-1], "区块1的数据", 6))
	blockChain = append(blockChain, blockchain.GenerateBlock(
		blockChain[len(blockChain)-1], "区块2的数据", 6))
	for _, block := range blockChain {
		slog.Info("区块信息",
			"索引", block.Index,
			"时间戳", block.Timestamp,
			"数据", block.Data,
			"前一个哈希值", block.PrevHash,
			"当前哈希值", block.Hash,
			"nonce", block.Nonce)
	}

}
