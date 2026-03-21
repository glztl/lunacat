package core

// 创建创世区块（第一个区块，没有上一个区块）
func CreateGenesisBlock() *Block {
	return CreateBlock(0, "0", "Genesis Block - LunaCat Blockchain")
}