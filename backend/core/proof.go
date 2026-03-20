package core

import "math/big"

// 挖矿难度：要求哈希前多少位是0 (例如 4 就是前4位是0)
const targetBits = 4

// ProofOfWork 代表工作量证明
type ProofOfWork struct {
	block *Block // 要验证的区块
	target *big.Int // 目标难度 (哈希必须小于这个值)
}

// NewProofOfWork 创建一个新的 PoW 实例
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// 左移 256 - targetBits 位，设置目标难度
	target.Lsh(target, uint(256 - targetBits))

	return &ProofOfWork{b, target}
}