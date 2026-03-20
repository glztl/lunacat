package core

// Blockchain 代表 LunaCat 区块链
type Blockchain struct {
	Blocks []Block  // 用切片存储所有区块
}

// 创建一个新的区块链 自动包含创世区块
func NewBlockchain() *Blockchain {
	genesisBlock := CreateGenesisBlock()
	return &Blockchain{
		Blocks: []Block{genesisBlock},
	}
}

// 向区块链中添加一个新区块
func (bc *Blockchain) AddBlock(data string) {
	// 获取链上最后一个区块
	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	// 创建新区块
	newBlock := CreateBlock(lastBlock.Index+1, lastBlock.Hash, data)

	// 添加到链上
	bc.Blocks = append(bc.Blocks, newBlock)
}