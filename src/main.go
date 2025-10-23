package main

import (
	"fmt"
	"time"
)

func main() {
	// 创世区块
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "创世区块",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)

	// 新区块
	newBlock := generateBlock(genesisBlock, "第二块区块的数据")

	fmt.Printf("创世区块的哈希值: %s\n", genesisBlock.Hash)
	fmt.Printf("新区块的哈希值: %s\n", newBlock.Hash)
}
