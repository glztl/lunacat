package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

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

// prepareData 准备用于计算哈希的数据（拼接所有字段 + Nonce）
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	// Integer to ASCII，将整数转换为字符串
	data := strconv.Itoa(pow.block.Index) +
		pow.block.Timestamp +
		pow.block.Data + 
		pow.block.PrevHash + 
		strconv.Itoa(targetBits) + 
		strconv.Itoa(nonce)
	return []byte(data)
}

// Run 执行工作量证明（挖矿）：找到符合要求的 Nonce 和哈希
func (pow *ProofOfWork) Run() (int, string) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Println("🔨 开始挖矿...")
	for nonce < 1e9 {
		// 防止无限循环
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)	// 长度32字节的固定数组
		hashInt.SetBytes(hash[:])

		// 检查哈希是否小于目标难度
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("✅ 挖矿成功! Nonce %d, Hash: %x\n", nonce, hash)
			break
		} else {
			nonce++
		}
	}

	return nonce, hex.EncodeToString(hash[:])
}

// Validate 验证区块是否符合 PoW 要求
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)

	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}