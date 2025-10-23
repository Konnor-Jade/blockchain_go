package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block 区块结构体
//
// # Index 区块索引
//
// # Timestamp 时间戳
//
// # Data 数据
//
// # PrevHash 前一个区块的哈希值
//
// Hash 当前区块的哈希值
type Block struct {
	Index     int    // 区块索引
	Timestamp string // 时间戳
	Data      string // 数据
	PrevHash  string // 前一个区块的哈希值
	Hash      string // 当前区块的哈希值
}

// calculateHash 计算区块的哈希值
//
// 参数:
//   - block 区块
//
// 返回值:
//   - 区块的哈希值
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

// generateBlock 生成新的区块
//
// 生成新的区块，设置索引为前一个区块的索引加1，时间戳为当前时间，数据为传入的数据，前一个哈希值为前一个区块的哈希值，当前哈希值为新计算的哈希值。
// 参数:
//   - prevBlock 前一个区块
//   - data 数据
//
// 返回值:
//   - 新的区块
func generateBlock(prevBlock Block, data string) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}
