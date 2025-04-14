# TON API Golang 封装库 (toncenterzp)

这是一个用于与 The Open Network (TON) 区块链交互的 Golang 封装库。该库封装了 GetBlock.io 提供的 TON API，使开发者能够轻松地在 Golang 应用程序中与 TON 区块链进行交互。

## 功能特点

- 完整封装 TON API 端点
- 支持 API 密钥认证
- 详细的错误处理和错误代码
- 类型安全的请求和响应结构
- 实用工具函数
- 丰富的示例程序

## 安装

使用 Go 模块安装此库：

```bash
go get github.com/zhaopeng331/toncenterzp
```

## 快速开始

以下是一个简单的示例，展示如何使用此库获取 TON 区块链的主链信息：

```go
package main

import (
	"fmt"
	"log"

	"github.com/zhaopeng331/toncenterzp"
)

func main() {
	// 创建 TON API 客户端
	client := toncenterzp.NewClient("YOUR-API-KEY")

	// 获取主链信息
	masterchainInfo, err := client.GetMasterchainInfo()
	if err != nil {
		log.Fatalf("获取主链信息失败: %v", err)
	}

	// 打印主链信息
	fmt.Printf("主链区块高度: %d\n", masterchainInfo.Result.LastBlockID.SeqNo)
	fmt.Printf("主链区块哈希: %s\n", masterchainInfo.Result.LastBlockID.RootHash)
}
```

## API 文档

### 客户端初始化

```go
// 使用默认配置创建客户端
client := toncenterzp.NewClient("YOUR-API-KEY")

// 使用自定义配置创建客户端
client := toncenterzp.NewClientWithOptions("YOUR-API-KEY", "https://custom-url.com/", 60*time.Second)
```

### 地址相关 API

- `DetectAddress(address string) (*DetectAddressResponse, error)`
- `EstimateFee(req EstimateFeeRequest) (*EstimateFeeResponse, error)`
- `GetAddressBalance(address string) (*GetAddressBalanceResponse, error)`
- `GetAddressInformation(address string) (*GetAddressInformationResponse, error)`
- `GetAddressState(address string) (*GetAddressStateResponse, error)`
- `GetExtendedAddressInformation(address string) (*GetExtendedAddressInformationResponse, error)`
- `GetWalletInformation(address string) (*GetWalletInformationResponse, error)`

### 区块相关 API

- `GetBlockHeader(req GetBlockHeaderRequest) (*GetBlockHeaderResponse, error)`
- `GetBlockTransactions(req GetBlockTransactionsRequest) (*GetBlockTransactionsResponse, error)`
- `GetConsensusBlock(req *GetConsensusBlockRequest) (*GetConsensusBlockResponse, error)`
- `LookupBlock(req LookupBlockRequest) (*LookupBlockResponse, error)`

### 主链相关 API

- `GetMasterchainBlockSignatures(req GetMasterchainBlockSignaturesRequest) (*GetMasterchainBlockSignaturesResponse, error)`
- `GetMasterchainInfo() (*GetMasterchainInfoResponse, error)`
- `GetShardBlockProof(req GetShardBlockProofRequest) (*GetShardBlockProofResponse, error)`
- `GetTokenData(address string) (*GetTokenDataResponse, error)`

### 交易相关 API

- `GetTransactions(req GetTransactionsRequest) (*GetTransactionsResponse, error)`
- `TryLocateResultTx(req TryLocateResultTxRequest) (*TryLocateResultTxResponse, error)`
- `TryLocateSourceTx(req TryLocateSourceTxRequest) (*TryLocateSourceTxResponse, error)`
- `TryLocateTx(req TryLocateTxRequest) (*TryLocateTxResponse, error)`

### RPC 相关 API

- `JSONRPC(method string, params interface{}) (*JSONRPCResponse, error)`
- `RunGetMethod(req RunGetMethodRequest) (*RunGetMethodResponse, error)`

### 发送相关 API

- `SendBoc(req SendBocRequest) (*SendBocResponse, error)`
- `SendBocReturnHash(req SendBocReturnHashRequest) (*SendBocReturnHashResponse, error)`
- `SendQuery(req SendQueryRequest) (*SendQueryResponse, error)`
- `Shards(seqNo int) (*ShardsResponse, error)`
- `PackAddress(address string) (*PackAddressResponse, error)`
- `UnpackAddress(address string) (*UnpackAddressResponse, error)`

### 工具函数

- `IsValidAddress(address string) bool`
- `FormatNanoTON(nanoTON string) (string, error)`
- `ValidateAPIKey(apiKey string) error`

## 示例程序

在 `examples` 目录中提供了多个示例程序，展示了如何使用此库：

- `address_info`: 获取地址详细信息
- `advanced_usage`: 高级用法示例
- `block_info`: 获取区块信息
- `block_scanner`: 扫描 TON 区块
- `test_client`: 测试客户端基本功能

## 错误处理

库使用 `ErrorWithCode` 结构体提供详细的错误信息和错误代码：

```go
type ErrorWithCode struct {
	Code    int
	Message string
	Err     error
}
```

错误代码分为三类：

- 客户端错误 (1xxx)
- 服务器错误 (2xxx)
- TON 特定错误 (3xxx)

## 许可证

此库采用 MIT 许可证。

## 贡献

欢迎提交问题和拉取请求。

## 致谢

感谢 GetBlock.io 提供的 TON API 服务。
