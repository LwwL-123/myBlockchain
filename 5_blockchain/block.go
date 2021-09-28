package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block 定义区块结构
type Block struct {
	Timestamp    int64  //时间戳
	Data         []byte //数据
	PreBlockHash []byte //前一块的hash
	Hash         []byte //当前块hash
}



// SetHash 区块内部设置hash的方法
func (b *Block) SetHash() {
	// 将时间戳转换为[]byte, strconv.FormatInt用于将整型数据转换成指定进制并以字符串的形式返回
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// 将前块的hash、交易信息、时间戳联合到一起
	headers := bytes.Join([][]byte{b.PreBlockHash,b.Data,timestamp},[]byte{})
	// 计算本块的hash值
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock 区块创建，返回指针
func NewBlock(data string, prevBlockHash []byte) *Block {
	// 构造Blcok
	block := &Block{time.Now().Unix(),[]byte(data),prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock 创世区块创建
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block",[]byte{})
}