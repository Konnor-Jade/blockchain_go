package blockchain

import "testing"

// TestGenerateBlock 测试创建区块
func TestGenerateBlock(t *testing.T) {
	// 创建前一个块（创世块）
	prevBlock := Block{
		Index:     0,
		Timestamp: "2024-11-28",
		Transactions: []Transaction{
			{
				ID: "genesis_tx",
				Outputs: []TransactionOutput{
					{Value: 100, Recipient: "address1"},
				},
			},
		},
		Hash: "1234567890abcdef",
	}

	// 创建新区块
	newBlock := GenerateBlock(prevBlock, []Transaction{
		{
			ID: "tx1",
			Outputs: []TransactionOutput{
				{Value: 50, Recipient: "address2"},
			},
		},
	}, 2)

	// 检查新块索引的正确性
	if newBlock.Index != prevBlock.Index+1 {
		t.Errorf("预期的索引是%d，收到的是%d", prevBlock.Index+1, newBlock.Index)
	}

	// 检查前一个块的哈希值是否正确
	if newBlock.PrevHash != prevBlock.Hash {
		t.Errorf("预期的前一个块哈希值是%s，收到的是%s", prevBlock.Hash, newBlock.PrevHash)
	}

	// 检查交易是否已添加
	if len(newBlock.Transactions) != 1 {
		t.Errorf("预期添加1笔交易，收到%d笔交易", len(newBlock.Transactions))
	}

	// 检查交易数据
	if newBlock.Transactions[0].Outputs[0].Recipient != "address2" {
		t.Errorf("预期的接收地址是address2，收到的是%s", newBlock.Transactions[0].Outputs[0].Recipient)
	}
}

// TestCreateTransaction 测试创建交易
func TestCreateTransaction(t *testing.T) {
	bc := Blockchain{}

	// 创世块，初始资金100
	genesisBlock := GenerateBlock(Block{}, []Transaction{
		{
			Outputs: []TransactionOutput{
				{Value: 100, Recipient: "address1"},
			},
		},
	}, 2)
	bc.Blocks = append(bc.Blocks, genesisBlock)

	// Creating a transaction
	tx, err := CreateTransaction("address1", "address2", 50, &bc)
	if err != nil {
		t.Fatalf("创建交易出错：%v", err)
	}

	// Checking the transfer amount
	if tx.Outputs[0].Value != 50 {
		t.Errorf("预期的转账金额是50，收到的是%d", tx.Outputs[0].Value)
	}

	// Checking the delivery amount
	if tx.Outputs[1].Value != 50 {
		t.Errorf("预期的转账金额是50，收到的是%d", tx.Outputs[1].Value)
	}
}
