package core


import (
	"sync"
)

// Blockchain 代表 LunaCat 区块链
type Blockchain struct {
	Blocks []*Block  // 用切片存储所有区块
}

var (
	bcInstance *Blockchain 	// 单例实例
	once	sync.Once	// 保证只初始化一次
)

// GetBlockchainInstance 获取区块链单例实例 (API 专用)
func GetBlockchainInstance() *Blockchain {
	once.Do(func() {
		bcInstance = NewBlockchain()
	})
	return bcInstance
}


// 创建一个新的区块链 自动包含创世区块
func NewBlockchain() *Blockchain {
	// genesisBlock := CreateGenesisBlock()
	// return &Blockchain{
	// 	Blocks: []Block{genesisBlock},
	// }
	return &Blockchain {
		Blocks: []*Block{CreateGenesisBlock()},
	}
}

// 向区块链中添加一个新区块
func (bc *Blockchain) AddBlock(data string) *Block {
	// 获取链上最后一个区块
	// lastBlock := bc.Blocks[len(bc.Blocks)-1]

	// 创建新区块
	// newBlock := CreateBlock(lastBlock.Index+1, lastBlock.Hash, data)

	// 添加到链上
	// bc.Blocks = append(bc.Blocks, newBlock)

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(prevBlock.Index+1, prevBlock.Hash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}