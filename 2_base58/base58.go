package main

import (
	"fmt"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func main() {
	result := Base58Encode(100)
	fmt.Println(string(result))
}

// 翻转数组
func ReverseBytes(data []byte) {
	for i := 0; i < len(data)/2; i++ {
		tmp := data[i]
		data[i] = data[len(data)-1-i]
		data[len(data)-1-i] = tmp
	}
}

//base58
func Base58Encode(input int64) []byte {
	var result []byte
	x := big.NewInt(input)
	//计算除数
	base := big.NewInt(int64(len(b58Alphabet)))
	//获取bigInt类型的0
	zero := big.NewInt(0)
	//存储余数
	mod := &big.Int{}
	//被除数不为0，一只计算
	for x.Cmp(zero) != 0 {
		//求余计算，x商值，mod余数
		x.DivMod(x, base, mod)
		//取出编码
		result = append(result, b58Alphabet[mod.Int64()])
	}
	ReverseBytes(result)

	return result
}
