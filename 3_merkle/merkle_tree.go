package main

import (
	"crypto/sha256"
	"fmt"
)

// MerkleNode 默克尔树节点结构
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// MerkleTree 默克尔树结构
type MerkleTree struct {
	//只记录根节点
	RootNode *MerkleNode
}

func main() {
	data := [][]byte{
		[]byte("my"),
		[]byte("name"),
		[]byte("is"),
		[]byte("lzw"),
	}

	tree := NewMerkleTree(data)
	showMerkleTree(tree.RootNode)
}

func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}
	// 如果left和right为空,则代表此节点为包含数据的叶子节点
	if left == nil && right == nil {
		// 计算哈希
		hash := sha256.Sum256(data)
		// 将[32]byte转换为[]byte
		mNode.Data = hash[:]
	} else {
		//将左子树和右子树数据合在一起
		prevHashes := append(left.Data, right.Data...)
		//计算哈希
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}
	mNode.Left = left
	mNode.Right = right

	return &mNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	//确保必须为2的整数倍节点
	if len(data)%2 != 0 {
		n := len(data) / 2
		data = data[:2*n]
	}

	//构建叶子节点
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	//两层循环完成节点树型的构造
	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		// i=0时，叶子节点hash合并，i=1时，合并新的nodes
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}
	return &mTree
}

//层次遍历，打印各个节点
func showMerkleTree(root *MerkleNode) {
	n := 0
	//创建队列，将root节点加入
	var queue []*MerkleNode
	queue = append(queue, root)

	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			fmt.Printf("层数为:%d,左节点[%p],右节点[%p],数据(%x)\n", n, node.Left, node.Right, node.Data)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
		queue = queue[length:]
		n++
	}
}
