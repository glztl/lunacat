package core

import (
	"fmt"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	// 创建创世区块
	// genesisBlock := CreateGenesisBlock()
	// fmt.Println("创世区块: ")
	// fmt.Printf("Index: %d\n", genesisBlock.Index)
	// fmt.Printf("Timestamp: %s\n", genesisBlock.Timestamp)
	// fmt.Printf("Data: %s\n", genesisBlock.Data)
	// fmt.Printf("PrevHash: %s\n", genesisBlock.PrevHash)
	// fmt.Printf("Hash: %s\n", genesisBlock.Hash)

	// 创建第二个区块
	// block2 := CreateBlock(1, genesisBlock.Hash, "Send 1 LunaCat to UserA")
	// fmt.Println("\n第二个区块: ")
	// fmt.Printf("Index: %d\n", block2.Index)
	// fmt.Printf("Timestamp: %s\n", block2.Timestamp)
	// fmt.Printf("Data: %s\n", block2.Data)
	// fmt.Printf("PrevHash: %s\n", block2.PrevHash)
	// fmt.Printf("Hash: %s\n", block2.Hash)

	// 1. 初始化区块链
	bc := NewBlockchain()
	fmt.Println(" 初始化 Lunacat 区块链成功!")
	fmt.Println("当前区块链上区块数量: ", len(bc.Blocks))

	// 2. 添加第一个新区块
	bc.AddBlock("Send 5 Lunacat to Alice")
	fmt.Println("\n 添加第一个新区块成功!")
	fmt.Println("当前区块链上区块数量: " , len(bc.Blocks))

	// 3. 添加第二个新区块
	bc.AddBlock("Send 3 Lunacat to Bob")
	fmt.Println("\n 添加第二个新区块成功!")
	fmt.Println("当前区块链上区块数量: " , len(bc.Blocks))

	// 4. 打印整条链
	fmt.Println("\n 打印整条 Lunacat 区块链: ")
	for i, block := range bc.Blocks {
		fmt.Printf("\n ---- 区块 %d ---- \n", i)
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
	}
}