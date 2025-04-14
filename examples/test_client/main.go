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

	// 测试 GetMasterchainInfo 方法
	fmt.Println("测试 GetMasterchainInfo 方法...")
	masterchainInfo, err := client.GetMasterchainInfo()
	if err != nil {
		log.Fatalf("GetMasterchainInfo 失败: %v", err)
	}
	fmt.Printf("当前主链区块高度: %d\n", masterchainInfo.Result.LastBlockID.SeqNo)
	fmt.Printf("主链区块哈希: %s\n", masterchainInfo.Result.LastBlockID.RootHash)
	fmt.Println("GetMasterchainInfo 测试成功")
	fmt.Println()

	// 测试 DetectAddress 方法
	fmt.Println("测试 DetectAddress 方法...")
	testAddress := "EQCkR1cGmnsE45N4K0otPl5EnxnRakmGqeJUNua5fkWhales"
	addressInfo, err := client.DetectAddress(testAddress)
	if err != nil {
		log.Fatalf("DetectAddress 失败: %v", err)
	}
	fmt.Printf("地址类型: %s\n", addressInfo.Result.GivenType)
	fmt.Printf("原始形式: %s\n", addressInfo.Result.RawForm)
	fmt.Println("DetectAddress 测试成功")
	fmt.Println()

	// 测试 GetAddressBalance 方法
	fmt.Println("测试 GetAddressBalance 方法...")
	balance, err := client.GetAddressBalance(testAddress)
	if err != nil {
		log.Fatalf("GetAddressBalance 失败: %v", err)
	}
	fmt.Printf("地址余额: %s nanoTON\n", balance.Result)
	
	// 格式化 nanoTON 为 TON
	formattedBalance, err := toncenterzp.FormatNanoTON(balance.Result)
	if err != nil {
		log.Fatalf("格式化余额失败: %v", err)
	}
	fmt.Printf("格式化余额: %s TON\n", formattedBalance)
	fmt.Println("GetAddressBalance 测试成功")
	fmt.Println()

	// 测试 GetWalletInformation 方法
	fmt.Println("测试 GetWalletInformation 方法...")
	walletInfo, err := client.GetWalletInformation(testAddress)
	if err != nil {
		log.Fatalf("GetWalletInformation 失败: %v", err)
	}
	fmt.Printf("钱包类型: %s\n", walletInfo.Result.WalletType)
	fmt.Printf("序列号: %d\n", walletInfo.Result.SeqNo)
	fmt.Println("GetWalletInformation 测试成功")
	fmt.Println()

	fmt.Println("所有测试通过！")
}
