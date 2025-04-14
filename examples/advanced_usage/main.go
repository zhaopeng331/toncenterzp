package main

import (
	"fmt"
	"log"

	"github.com/zhaopeng331/toncenterzp"
)

func main() {
	// 替换为您的 API 密钥
	apiKey := "YOUR-API-KEY"

	// 创建 TON API 客户端
	client := toncenterzp.NewClient(apiKey)

	// 示例地址
	tokenAddress := "EQBynBO23ywHy_CgarY9NK9FTz0yDsG82PtcbSTQgGoXwiuA" // 示例 Jetton 地址

	// 获取代币信息
	fmt.Println("获取代币信息...")
	tokenData, err := client.GetTokenData(tokenAddress)
	if err != nil {
		log.Fatalf("获取代币信息失败: %v", err)
	}

	// 打印代币信息
	fmt.Printf("代币名称: %s\n", tokenData.Result.Name)
	fmt.Printf("代币符号: %s\n", tokenData.Result.Symbol)
	fmt.Printf("代币精度: %d\n", tokenData.Result.Decimals)
	fmt.Printf("代币地址: %s\n", tokenData.Result.Address)
	fmt.Println()

	// 示例地址
	address := "EQCkR1cGmnsE45N4K0otPl5EnxnRakmGqeJUNua5fkWhales"

	// 运行合约方法
	fmt.Println("运行合约方法...")
	runMethodResult, err := client.RunGetMethod(toncenterzp.RunGetMethodRequest{
		Address: address,
		Method:  "seqno", // 获取钱包序列号
		Stack:   []interface{}{},
	})
	if err != nil {
		log.Printf("运行合约方法失败: %v", err)
		fmt.Println("注意: 某些地址可能不支持此方法，这是正常的")
	} else {
		// 打印方法执行结果
		fmt.Printf("方法执行结果:\n")
		fmt.Printf("  退出码: %d\n", runMethodResult.Result.ExitCode)
		fmt.Printf("  使用的 gas: %d\n", runMethodResult.Result.GasUsed)
		fmt.Printf("  堆栈结果: %v\n", runMethodResult.Result.Stack)
	}
	fmt.Println()

	// 使用 JSON-RPC 接口
	fmt.Println("使用 JSON-RPC 接口...")
	rpcParams := map[string]interface{}{
		"address": address,
	}
	rpcResult, err := client.JSONRPC("getWalletInformation", rpcParams)
	if err != nil {
		log.Fatalf("JSON-RPC 调用失败: %v", err)
	}

	// 打印 JSON-RPC 结果
	fmt.Printf("JSON-RPC 结果: %s\n", string(rpcResult.Result))
	fmt.Println()

	// 解析地址
	fmt.Println("解析地址...")
	unpackResult, err := client.UnpackAddress(address)
	if err != nil {
		log.Fatalf("解析地址失败: %v", err)
	}

	// 打印地址解析结果
	fmt.Printf("原始形式: %s\n", unpackResult.Result.RawForm)
	fmt.Printf("工作链: %d\n", unpackResult.Result.WorkChain)
	fmt.Printf("哈希部分: %s\n", unpackResult.Result.Hash)
	fmt.Printf("是否可弹跳: %t\n", unpackResult.Result.Bounceable)
	fmt.Printf("是否为测试网络: %t\n", unpackResult.Result.TestOnly)
	fmt.Println()

	fmt.Println("示例完成!")
}
