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

	// 获取主链信息
	fmt.Println("获取主链信息...")
	masterchainInfo, err := client.GetMasterchainInfo()
	if err != nil {
		log.Fatalf("获取主链信息失败: %v", err)
	}

	// 打印主链信息
	fmt.Printf("主链区块高度: %d\n", masterchainInfo.Result.LastBlockID.SeqNo)
	fmt.Printf("主链区块哈希: %s\n", masterchainInfo.Result.LastBlockID.RootHash)
	fmt.Printf("主链分片: %s\n", masterchainInfo.Result.LastBlockID.Shard)
	fmt.Println()

	// 获取分片信息
	fmt.Println("获取分片信息...")
	shards, err := client.Shards(masterchainInfo.Result.LastBlockID.SeqNo)
	if err != nil {
		log.Fatalf("获取分片信息失败: %v", err)
	}

	// 打印分片信息
	fmt.Printf("分片数量: %d\n", len(shards.Result.Shards))
	for i, shard := range shards.Result.Shards {
		fmt.Printf("分片 #%d:\n", i+1)
		fmt.Printf("  工作链: %d\n", shard.Workchain)
		fmt.Printf("  分片ID: %s\n", shard.Shard)
		fmt.Printf("  区块高度: %d\n", shard.SeqNo)
		fmt.Printf("  根哈希: %s\n", shard.RootHash)
	}
	fmt.Println()

	// 获取区块头信息
	fmt.Println("获取区块头信息...")
	blockHeader, err := client.GetBlockHeader(toncenterzp.GetBlockHeaderRequest{
		Workchain: masterchainInfo.Result.LastBlockID.Workchain,
		Shard:     masterchainInfo.Result.LastBlockID.Shard,
		SeqNo:     masterchainInfo.Result.LastBlockID.SeqNo,
	})
	if err != nil {
		log.Fatalf("获取区块头信息失败: %v", err)
	}

	// 打印区块头信息
	fmt.Printf("区块全局ID: %d\n", blockHeader.Result.GlobalID)
	fmt.Printf("区块版本: %d\n", blockHeader.Result.Version)
	fmt.Printf("生成时间: %d\n", blockHeader.Result.GenUtime)
	fmt.Printf("是否为关键区块: %t\n", blockHeader.Result.IsKeyBlock)
	fmt.Printf("开始逻辑时间: %s\n", blockHeader.Result.StartLt)
	fmt.Printf("结束逻辑时间: %s\n", blockHeader.Result.EndLt)
	fmt.Println()

	// 获取区块交易
	fmt.Println("获取区块交易...")
	blockTransactions, err := client.GetBlockTransactions(toncenterzp.GetBlockTransactionsRequest{
		Workchain: masterchainInfo.Result.LastBlockID.Workchain,
		Shard:     masterchainInfo.Result.LastBlockID.Shard,
		SeqNo:     masterchainInfo.Result.LastBlockID.SeqNo,
		Count:     10, // 只获取最多10笔交易
	})
	if err != nil {
		log.Fatalf("获取区块交易失败: %v", err)
	}

	// 打印区块交易信息
	fmt.Printf("交易数量: %d\n", len(blockTransactions.Result.Transactions))
	fmt.Printf("是否不完整: %t\n", blockTransactions.Result.Incomplete)
	for i, tx := range blockTransactions.Result.Transactions {
		fmt.Printf("交易 #%d:\n", i+1)
		fmt.Printf("  账户: %s\n", tx.Account)
		fmt.Printf("  哈希: %s\n", tx.Hash)
		fmt.Printf("  逻辑时间: %s\n", tx.Lt)
	}
}
