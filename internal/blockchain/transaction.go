package blockchain

type TransactionInput struct {
	PrevTxID  string `json:"prevTxID"`    // 前一个交易的ID
	OutputIdx int    `json:"outputIndex"` // 前一个交易输出的索引
	Signature string `json:"signature"`   // 签名确认有权使用资金
}

type TransactionOutput struct {
	Value     int    `json:"value"`     // 输出的金额
	Recipient string `json:"recipient"` // 接收者的地址
}

type Transaction struct {
	ID       string              `json:"id"`       // 交易的唯一标识符
	Inputs   []TransactionInput  `json:"inputs"`   // 输入的交易列表
	Outputs  []TransactionOutput `json:"outputs"`  // 输出的交易列表
	Coinbase bool                `json:"coinbase"` // 是否为创世交易
}
