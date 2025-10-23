package main

import (
	"context"
	"log/slog"
)

func main() {
	initBlockchain()

	addBlock("区块1的数据")
	addBlock("区块2的数据")

	for _, block := range Blockchain {
		slog.Log(
			context.Background(),
			slog.LevelInfo,
			"区块信息",
			slog.Int("索引", block.Index),
			slog.String("时间戳", block.Timestamp),
			slog.String("数据", block.Data),
			slog.String("前一个哈希值", block.PrevHash),
			slog.String("当前哈希值", block.Hash),
		)
	}
}
