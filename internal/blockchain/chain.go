package blockchain

import (
	"blockchain_go/internal/storage"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
)

// Blockchain 区块链结构体
//
// 字段：
//   - Blocks 区块链中的所有区块
//   - Storage 用于存储区块数据的存储接口
//
// 方法：
//   - LoadBlockchain 加载区块链
type Blockchain struct {
	Blocks  []Block
	Storage storage.Storage
}

// AddBlock 添加新区块
//
// 参数：
//   - data 新区块的数据
//
// 功能：
//   - 生成一个新的区块
//   - 将新的区块添加到区块链中
func (bc *Blockchain) AddBlock(block Block) error {
	bc.Blocks = append(bc.Blocks, block)

	// 保存区块到内存存储中
	blockData, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return bc.Storage.SaveBlock(block.Hash, blockData)
}

// LoadBlockchain 加载区块链
//
// 功能：
//   - 从存储中加载所有区块
//   - 将加载的区块添加到区块链中
func (bc *Blockchain) LoadBlockchain() error {
	err := bc.Storage.Iterate(func(hash string, data []byte) error {
		var block Block
		if err := json.Unmarshal(data, &block); err != nil {
			return err
		}
		bc.Blocks = append(bc.Blocks, block)
		return nil
	})
	if err != nil {
		return err
	}
	slog.Info("成功加载区块链", "区块数量", len(bc.Blocks))
	return nil
}

// GetBalance 获取地址余额
//
// 参数：
//   - address 地址
//
// 返回：
//   - balance 地址余额
func (bc *Blockchain) GetBalance(address string) int {
	balance := 0
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			for _, output := range tx.Outputs {
				if output.Recipient == address {
					balance += output.Value
				}
			}
		}
	}
	return balance
}

// CreateTransaction 创建交易
//
// 参数：
//   - sender 发送者地址
//   - recipient 接收者地址
//   - amount 交易金额
//   - bc 区块链实例
//
// 返回：
//   - tx 交易实例
//   - err 错误信息
func CreateTransaction(sender, recipient string, amount int, bc *Blockchain) (*Transaction, error) {
	// 检查发送者余额是否足够
	balance := bc.GetBalance(sender)
	if balance < amount {
		return nil, errors.New("发送者余额不足")
	}

	// 创建交易
	tx := &Transaction{
		Inputs: []TransactionInput{
			{PrevTxID: "", OutputIdx: 0, Signature: sender},
		},
		Outputs: []TransactionOutput{
			{Value: balance - amount, Recipient: recipient},
		},
	}
	tx.ID = CalculateTransactionID(tx)
	return tx, nil
}

// CalculateTransactionID 计算交易ID
//
// 参数：
//   - tx 交易实例
//
// 返回：
//   - id 交易ID
func CalculateTransactionID(tx *Transaction) string {
	data := fmt.Sprintf("%v", tx)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// CreateCoinbaseTransaction 创建挖矿奖励交易
//
// 参数：
//   - recipient 接收者地址
//   - amount 交易金额
//
// 返回：
//   - tx 创世交易实例
//
// 功能：
//   - 创建一个创世交易，用于奖励矿工
func CreateCoinbaseTransaction(recipient string, amount int) *Transaction {
	tx := &Transaction{
		Inputs:   []TransactionInput{},
		Outputs:  []TransactionOutput{{Value: amount, Recipient: recipient}},
		Coinbase: true,
	}
	tx.ID = CalculateTransactionID(tx)
	return tx
}
