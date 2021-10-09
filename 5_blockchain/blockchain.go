package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blockBucket = "blocks"

// BlockChain 区块链
type BlockChain struct {
	// 用来记录区块的hash
	tip []byte
	db  *bolt.DB
}

// BlockchainIterator 迭代器
type BlockchainIterator struct {

}

// AddBlock 增加区块
func (bc *BlockChain) AddBlock(data string) {
	var tip []byte
	// 1. 获取tip值，此时不能再打开数据库文件，要用区块的结构
	bc.db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(blockBucket))
		tip = buck.Get([]byte("last"))
		return nil
	})

	// 2. 更新数据库
	bc.db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte(blockBucket))
		block := NewBlock(data,tip)
		// 将新区块放入db
		buck.Put(block.Hash,block.Serialize())
		buck.Put([]byte("last"),block.Hash)
		// 覆盖tip
		bc.tip = block.Hash
		return nil
	})
}

func NewBlockChain() *BlockChain {
	var tip []byte
	// 1. 打开数据库文件
	db, _ := bolt.Open(dbFile, 0600, nil)
	// 2. 更新数据库
	db.Update(func(tx *bolt.Tx) error {
		// 2.1 获取bucket
		buck := tx.Bucket([]byte(blockBucket))
		if buck == nil {
			//2.2.1 第一次使用 创建创世块
			fmt.Printf("没有存在的区块链，创建一个新的区块链")
			genesis := NewGenesisBlock()
			//2.2.2 区块数据编码
			block_data := genesis.Serialize()
			//2.2.3 创建新的bucket，存入区块信息
			bucket, _ := tx.CreateBucket([]byte(blockBucket))
			bucket.Put(genesis.Hash, block_data)
			bucket.Put([]byte("last"), genesis.Hash)
			tip = genesis.Hash
		} else {
			//2.3 不是创世区块
			tip = buck.Get([]byte("last"))
		}
		return nil
	})

	// 3. 记录Blockchain信息
	return &BlockChain{tip,db}
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("1111")
	bc.AddBlock("2222")
	//bc.AddBlock("3333")
	//bc.AddBlock("4444")

	//区块链遍历
	for _, block := range bc.blocks {
		fmt.Printf("前一块的hash:%x\n", block.PreBlockHash)
		fmt.Printf("数据：%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("Nonce:%d\n", block.Nonce)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %t\n", pow.Validate())
		fmt.Println()
	}
}
