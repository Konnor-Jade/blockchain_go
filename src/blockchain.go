package main

import "time"

// Blockchain 用于存储整个区块链
var Blockchain []Block

// createGenesisBlock 创建创世区块
//
// 返回：
//   - genesisBlock 创世区块
func createGenesisBlock() Block {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "创世区块",
		PrevHash:  "",
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}

// initBlockchain 初始化区块链
//
// 功能：
//   - 创建创世区块
//   - 将创世区块添加到区块链中
func initBlockchain() {
	genesisBlock := createGenesisBlock()
	Blockchain = append(Blockchain, genesisBlock)
}

// addBlock 添加新区块
//
// 参数：
//   - data 新区块的数据
//
// 功能：
//   - 生成一个新的区块
//   - 将新的区块添加到区块链中
func addBlock(data string) {
	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := generateBlock(prevBlock, data)
	Blockchain = append(Blockchain, newBlock)
}
