package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"strconv"
	"strings"
)

// ProofOfWork 工作量证明结构体
type ProofOfWork struct {
	Block  *Block
	Target string // 目标哈希值中的前导0的数量
}

// NewProofOfWork 创建一个新的工作量证明实例
//
// 参数：
//   - block 待证明的区块
//   - difficulty 目标哈希值中的前导0的数量
//
// 返回：
//   - pow 工作量证明实例
func NewProofOfWork(block *Block, difficulty int) *ProofOfWork {
	target := strings.Repeat("0", difficulty)
	return &ProofOfWork{
		Block:  block,
		Target: target,
	}
}

// Run 执行工作量证明
//
// 返回：
//   - hash 符合目标哈希值的哈希值
//   - nonce 找到符合目标哈希值的nonce值
func (pow *ProofOfWork) Run() (string, int) {
	var nonce int
	var hash string
	for {
		data := pow.Block.Data + strconv.Itoa(pow.Block.Index) + pow.Block.PrevHash + strconv.Itoa(nonce)
		hashBytes := sha256.Sum256([]byte(data))
		hash = hex.EncodeToString(hashBytes[:])
		if strings.HasPrefix(hash, pow.Target) {
			break
		}
		nonce++
	}
	slog.Info("工作量证明成功", "nonce", nonce, "hash", hash)
	return hash, nonce
}
