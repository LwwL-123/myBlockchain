package main

import "fmt"

// BlockChain 区块链
type BlockChain struct {
	blocks []*Block
}

// AddBlock 增加区块
func (bc *BlockChain) AddBlock(data string) {
	// 获取前块信息
	prevBlock := bc.blocks[len(bc.blocks)-1]
	// 利用前块生成新区块
	newBlock := NewBlock(data, prevBlock.Hash)
	// 添加到区块链中
	bc.blocks = append(bc.blocks,newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("1111")
	bc.AddBlock("2222")
	//bc.AddBlock("3333")
	//bc.AddBlock("4444")

	//区块链遍历
	for _, block := range bc.blocks{
		fmt.Printf("前一块的hash:%x\n",block.PreBlockHash)
		fmt.Printf("数据：%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("Nonce:%d\n",block.Nonce)
		pow := NewProofOfWork(block)
		fmt.Printf("Pow: %t\n",pow.Validate())
		fmt.Println()
	}
}
