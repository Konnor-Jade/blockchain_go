# blockchain_go

## 项目简介

`blockchain_go` 是一个使用 Go 语言实现的轻量级区块链演示项目，旨在展示区块链技术的核心概念和工作原理。

## 关键特性

- 🔒 工作量证明（Proof of Work）共识机制
- 🔑 地址生成与管理
- 💰 简单的交易系统
- 💾 BoltDB 持久化存储
- 🖥️ 命令行交互界面

## 技术亮点

### 工作量证明（PoW）实现

项目通过动态难度控制实现了简单但有效的工作量证明机制：

```go

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

    return hash, nonce
}


```

- 通过前导零数量控制挖矿难度
- 使用 SHA-256 哈希算法
- 动态调整 Nonce 值直到找到符合条件的哈希

### 项目结构

```
blockchain_go/
├── cmd/main/ # 程序入口
├── internal/
│ ├── blockchain/ # 核心区块链逻辑
│ ├── storage/ # 存储接口
│ └── utils/ # 工具函数
```

## 功能列表

1. 地址生成 `gen-address`
2. 余额充值 `fund`
3. 余额查询 `balance`
4. 交易转账 `send`
5. 区块链浏览 `view`

## 快速开始

```bash
# 克隆项目
git clone https://github.com/yourusername/blockchain_go.git

# 运行项目
go run cmd/main/main.go
```

## 交互示例

```
> gen-address
您的新地址: acaccdf21bd7ffaf5f43df2580954a1ff87c54bec9a4ba020cbb033069d78223
> fund
接收地址: acaccdf21bd7ffaf5f43df2580954a1ff87c54bec9a4ba020cbb033069d78223
金额: 100
```
