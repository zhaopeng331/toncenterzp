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
	address := "EQCkR1cGmnsE45N4K0otPl5EnxnRakmGqeJUNua5fkWhales"

	// 获取地址信息
	fmt.Println("获取地址信息...")
	addressInfo, err := client.GetAddressInformation(address)
	if err != nil {
		log.Fatalf("获取地址信息失败: %v", err)
	}

	// 打印地址信息
	fmt.Printf("地址: %s\n", addressInfo.Result.Address)
	fmt.Printf("余额: %s nanoTON\n", addressInfo.Result.Balance)
	fmt.Printf("状态: %s\n", addressInfo.Result.State)
	fmt.Printf("账户状态: %s\n", addressInfo.Result.AccountStatus)
	fmt.Println()

	// 获取钱包信息
	fmt.Println("获取钱包信息...")
	walletInfo, err := client.GetWalletInformation(address)
	if err != nil {
		log.Fatalf("获取钱包信息失败: %v", err)
	}

	// 打印钱包信息
	fmt.Printf("是否为钱包: %t\n", walletInfo.Result.Wallet)
	fmt.Printf("钱包类型: %s\n", walletInfo.Result.WalletType)
	fmt.Printf("序列号: %d\n", walletInfo.Result.SeqNo)
	fmt.Printf("钱包ID: %d\n", walletInfo.Result.WalletID)
	fmt.Printf("公钥: %s\n", walletInfo.Result.PublicKey)
	fmt.Println()

	// 获取交易历史
	fmt.Println("获取交易历史...")
	transactions, err := client.GetTransactions(toncenterzp.GetTransactionsRequest{
		Address: address,
		Limit:   5, // 只获取最近5笔交易
	})
	if err != nil {
		log.Fatalf("获取交易历史失败: %v", err)
	}

	// 打印交易信息
	fmt.Printf("最近 %d 笔交易:\n", len(transactions.Result.Transactions))
	for i, tx := range transactions.Result.Transactions {
		fmt.Printf("交易 #%d:\n", i+1)
		fmt.Printf("  哈希: %s\n", tx.Hash)
		fmt.Printf("  时间: %d\n", tx.Now)
		fmt.Printf("  费用: %s nanoTON\n", tx.TotalFees)
		
		// 如果有输入消息，打印输入消息信息
		if tx.InMsg.Source != "" {
			fmt.Printf("  输入消息:\n")
			fmt.Printf("    来源: %s\n", tx.InMsg.Source)
			fmt.Printf("    金额: %s nanoTON\n", tx.InMsg.Value)
			if tx.InMsg.MsgData.Text != "" {
				fmt.Printf("    消息: %s\n", tx.InMsg.MsgData.Text)
			}
		}
		
		// 如果有输出消息，打印输出消息信息
		if len(tx.OutMsgs) > 0 {
			fmt.Printf("  输出消息数量: %d\n", len(tx.OutMsgs))
			for j, outMsg := range tx.OutMsgs {
				fmt.Printf("  输出消息 #%d:\n", j+1)
				fmt.Printf("    目标: %s\n", outMsg.Destination)
				fmt.Printf("    金额: %s nanoTON\n", outMsg.Value)
				if outMsg.MsgData.Text != "" {
					fmt.Printf("    消息: %s\n", outMsg.MsgData.Text)
				}
			}
		}
		fmt.Println()
	}
}
