package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

var (
	// Nonce循环上限
	maxNonce = math.MaxInt64
)

// 难度值
const targetBits = 24

// ProofOfWork Pow结构
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	//target左移256-24（挖矿难度）
	target.Lsh(target, uint(256-targetBits))
	//生成pow结构
	pow := &ProofOfWork{
		block:  b,
		target: target,
	}
	return pow
}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	var nonce int64 = 0

	fmt.Printf("Mining the block containning %s, maxNonce = %d\n", pow.block.Data, maxNonce)
	for nonce < int64(maxNonce) {
		// 数据准备
		data := pow.prepareDate(nonce)
		// 计算hash
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])

		// 按字节进行比较，hashInt.Cmp小于0代表找到目标nonce
		if hashInt.Cmp(pow.target) == -1 {
			break
		}else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce,hash[:]
}

// 准备数据
func (pow *ProofOfWork) prepareDate(nonce int64) []byte {
	//Join将s的元素连接起来以创建一个新的字节片。分隔符sep放置在所得切片中的元素之间。func Join(s [][]byte, sep []byte) []byte
	data := bytes.Join(
		[][]byte{
			pow.block.PreBlockHash,
			pow.block.Data,
			Int2Hex(pow.block.Timestamp),
			Int2Hex(int64(targetBits)),
			Int2Hex(nonce),
		},
		[]byte{ },
	)
	return data
}

// Int2Hex 将int转换为byte
func Int2Hex(num int64) []byte {
	// 创建缓冲区
	buff := new(bytes.Buffer)
	//大端法写入
	binary.Write(buff,binary.BigEndian,num)
	// 返回缓冲区的字节切片
	return buff.Bytes()
}

// Validate 检验区块的正确性
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareDate(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}