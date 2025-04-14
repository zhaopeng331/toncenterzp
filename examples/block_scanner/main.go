package main

import (
	"fmt"
	"log"
	"time"

	"github.com/zhaopeng331/toncenterzp"
)

func main() {
	// 替换为您的 API 密钥
	apiKey := "YOUR-API-KEY"

	// 创建 TON API 客户端
	client := toncenterzp.NewClient(apiKey)

	// 记录上次扫描的区块高度
	var lastSeqNo int
	var currentSeqNo int

	// 首次获取主链信息
	masterchainInfo, err := client.GetMasterchainInfo()
	if err != nil {
		log.Fatalf("获取主链信息失败: %v", err)
	}

	// 设置初始区块高度
	lastSeqNo = masterchainInfo.Result.LastBlockID.SeqNo
	fmt.Printf("初始区块高度: %d\n", lastSeqNo)

	// 持续扫描区块
	for {
		// 获取最新的主链信息
		masterchainInfo, err := client.GetMasterchainInfo()
		if err != nil {
			log.Printf("获取主链信息失败: %v", err)
			time.Sleep(5 * time.Second) // 出错时等待5秒后重试
			continue
		}

		// 获取当前区块高度
		currentSeqNo = masterchainInfo.Result.LastBlockID.SeqNo

		// 如果区块高度没有变化，等待一分钟后继续
		if currentSeqNo <= lastSeqNo {
			fmt.Printf("当前区块高度 %d 未变化，等待一分钟后继续...\n", currentSeqNo)
			time.Sleep(1 * time.Minute)
			continue
		}

		// 处理新区块
		for seqNo := lastSeqNo + 1; seqNo <= currentSeqNo; seqNo++ {
			fmt.Printf("处理区块 #%d\n", seqNo)

			// 获取区块交易
			blockTransactions, err := client.GetBlockTransactions(toncenterzp.GetBlockTransactionsRequest{
				Workchain: masterchainInfo.Result.LastBlockID.Workchain,
				Shard:     masterchainInfo.Result.LastBlockID.Shard,
				SeqNo:     seqNo,
			})

			if err != nil {
				log.Printf("获取区块 #%d 交易失败: %v", seqNo, err)
				continue
			}

			// 处理区块中的交易
			fmt.Printf("区块 #%d 包含 %d 笔交易\n", seqNo, len(blockTransactions.Result.Transactions))
			for _, tx := range blockTransactions.Result.Transactions {
				fmt.Printf("  交易哈希: %s, 账户: %s\n", tx.Hash, tx.Account)
				
				// 这里可以添加更多的交易处理逻辑
				// 例如：检查特定地址的交易、分析交易金额等
			}
		}

		// 更新上次扫描的区块高度
		lastSeqNo = currentSeqNo
		fmt.Printf("已更新到区块高度: %d\n", lastSeqNo)

		// 短暂休息，避免频繁请求
		time.Sleep(5 * time.Second)
	}
}
