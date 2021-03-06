package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

// Block 定义区块结构
type Block struct {
	Timestamp    int64  //时间戳
	Data         []byte //数据
	PreBlockHash []byte //前一块的hash
	Hash         []byte //当前块hash
	Nonce        int64  // 随机值
}

// SetHash 区块内部设置hash的方法
func (b *Block) SetHash() {
	// 将时间戳转换为[]byte, strconv.FormatInt用于将整型数据转换成指定进制并以字符串的形式返回
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// 将前块的hash、交易信息、时间戳联合到一起
	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	// 计算本块的hash值
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlock 区块创建，返回指针
func NewBlock(data string, prevBlockHash []byte) *Block {
	// 构造Blcok
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{},0}
	//block.SetHash()

	// 先挖矿
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	// 设置hash和nonce
	block.Hash = hash
	block.Nonce = nonce
	return block
}

// NewGenesisBlock 创世区块创建
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Serialize 序列化区块
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	// 编码器
	encoder := gob.NewEncoder(&result)
	// 编码
	encoder.Encode(b)
	return result.Bytes()
}

// DeSerialize 区块数据还原为block
func DeSerialize(d []byte) *Block {
	var block Block
	// 创建解码器
	decoder := gob.NewDecoder(bytes.NewReader(d))
	// 解析区块数据
	decoder.Decode(&block)
	return &block
}