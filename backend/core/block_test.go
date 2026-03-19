package core

import (
	"fmt"
	"testing"
)

func TestCreateBlock(t *testing.T) {
	// 创建创世区块
	genesisBlock := CreateGenesisBlock()
	fmt.Println("创世区块: ")
	fmt.Printf("Index: %d\n", genesisBlock.Index)
	fmt.Printf("Timestamp: %s\n", genesisBlock.Timestamp)
	fmt.Printf("Data: %s\n", genesisBlock.Data)
	fmt.Printf("PrevHash: %s\n", genesisBlock.PrevHash)
	fmt.Printf("Hash: %s\n", genesisBlock.Hash)

	// 创建第二个区块
	block2 := CreateBlock(1, genesisBlock.Hash, "Send 1 LunaCat to UserA")
	fmt.Println("\n第二个区块: ")
	fmt.Printf("Index: %d\n", block2.Index)
	fmt.Printf("Timestamp: %s\n", block2.Timestamp)
	fmt.Printf("Data: %s\n", block2.Data)
	fmt.Printf("PrevHash: %s\n", block2.PrevHash)
	fmt.Printf("Hash: %s\n", block2.Hash)
}