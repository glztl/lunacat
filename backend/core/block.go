package core

import (
	"time"
)

// Block
type Block struct {
	Index     int    `json:"index"`     // 区块高度（编号）
	Timestamp string `json:"timestamp"` // 时间戳
	Data      string `json:"data"`      // 区块存储的数据（交易）
	PrevHash  string `json:"prevHash"`  // 上一个区块的哈希
	Hash      string `json:"hash"`      // 当前区块的哈希

	Nonce	  int     `json:"nonce"`     // 工作量证明的随机数
}

// 计算区块的哈希值
func calculateHash(block Block) string {
	// 拼接区块所有字段作为哈希源数据
	// record := strconv.Itoa(block.Index) +
	// 	block.Timestamp +
	// 	block.Data +
	// 	block.PrevHash

	// // 计算 SHA256 哈希
	// h := sha256.New()
	// h.Write([]byte(record))
	// hashed := h.Sum(nil)

	// // 返回十六进制字符串
	// return hex.EncodeToString(hashed)

	return ""
}

// 创建一个新区块，并自动计算哈希
func CreateBlock(index int, prevHash string, data string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
	}

	// 使用PoW挖矿，获取Nonce和Hash
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()

	// 计算当前区块哈希
	block.Nonce = nonce
	block.Hash = hash
	return block
}
