package main

import "testing"

func Test_initBlockchain(t *testing.T) {
	initBlockchain()

	addBlock("区块1的数据")
	addBlock("区块2的数据")

	for _, block := range Blockchain {
		t.Logf("区块索引：%d，时间戳：%s，数据：%s，前一个哈希值：%s，当前哈希值：%s\n",
			block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash)
	}

}
