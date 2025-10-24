package main

import (
	"blockchain_go/internal/blockchain"
	"blockchain_go/internal/storage"
	"blockchain_go/internal/utils"
	"fmt"
	"log"
	"time"
)

func main() {
	// 打开数据库，并读取
	db, err := storage.NewBoltDB("blockchain.db")
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer db.Close()

	bc := &blockchain.Blockchain{Storage: db}
	if err := bc.LoadBlockchain(); err != nil {
		log.Fatalf("加载区块链失败: %v", err)
	}

	// CLI interface
	var command string
	for {
		fmt.Println("请输入命令 (gen-address, fund, balance, send, view, exit):")
		fmt.Scanln(&command)

		switch command {
		case "gen-address":
			address := utils.GenerateAddress()
			fmt.Println("您的新地址:", address)

		case "fund":
			var recipient string
			var amount int
			fmt.Println("请输入接收地址:")
			fmt.Scanln(&recipient)
			fmt.Println("请输入金额:")
			fmt.Scanln(&amount)

			// 创建创世交易
			coinbaseTx := blockchain.CreateCoinbaseTransaction(recipient, amount)

			var newBlock blockchain.Block
			if len(bc.Blocks) == 0 {
				// 如果区块链为空，创建创世区块
				newBlock = blockchain.Block{
					Index:        0,
					Timestamp:    time.Now().Format(time.RFC3339),
					Transactions: []blockchain.Transaction{*coinbaseTx},
					PrevHash:     "0",
					Hash:         "",
					Nonce:        0,
				}
				// 计算创世区块的哈希
				pow := blockchain.NewProofOfWork(&newBlock, 4)
				hash, nonce := pow.Run()
				newBlock.Hash = hash
				newBlock.Nonce = nonce
			} else {
				// 如果区块链不为空，使用原来的方法
				newBlock = blockchain.GenerateBlock(bc.Blocks[len(bc.Blocks)-1], []blockchain.Transaction{*coinbaseTx}, 4)
			}

			if err := bc.AddBlock(newBlock); err != nil {
				log.Printf("添加创世交易到新区块失败: %v", err)
			} else {
				fmt.Println("余额已补充!")
			}

		case "balance":
			var address string
			fmt.Println("请输入地址:")
			fmt.Scanln(&address)
			fmt.Printf("地址 %s 的余额: %d\n", address, bc.GetBalance(address))

		case "send":
			var sender, recipient string
			var amount int
			fmt.Println("请输入发送地址:")
			fmt.Scanln(&sender)
			fmt.Println("请输入接收地址:")
			fmt.Scanln(&recipient)
			fmt.Println("请输入金额:")
			fmt.Scanln(&amount)

			tx, err := blockchain.CreateTransaction(sender, recipient, amount, bc)
			if err != nil {
				fmt.Println("创建交易失败:", err)
				continue
			}

			// 添加交易到新区块
			newBlock := blockchain.GenerateBlock(bc.Blocks[len(bc.Blocks)-1], []blockchain.Transaction{*tx}, 4)
			if err := bc.AddBlock(newBlock); err != nil {
				fmt.Printf("添加交易到新区块失败: %v\n", err)
			} else {
				fmt.Println("交易已添加到新区块!")
			}

		case "view":
			for _, block := range bc.Blocks {
				fmt.Printf("区块 %d: %s\n", block.Index, block.Hash)
				for _, tx := range block.Transactions {
					fmt.Printf("  交易 %s: %d -> %s\n", tx.ID, tx.Outputs[0].Value, tx.Outputs[0].Recipient)
				}
			}

		case "exit":
			return

		default:
			fmt.Println("未知命令.")
		}
	}
}
