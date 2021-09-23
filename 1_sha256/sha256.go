package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str := []byte("lzw")
	// 计算hash值
	hash := sha256.Sum256(str)
	fmt.Println(hash)
}
