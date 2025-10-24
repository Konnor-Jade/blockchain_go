package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
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
	Index        int           // 区块索引
	Timestamp    string        // 时间戳
	Data         string        // 数据
	PrevHash     string        // 前一个区块的哈希值
	Hash         string        // 当前区块的哈希值
	Nonce        int           // 工作量证明的nonce值
	Transactions []Transaction // 交易列表
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
