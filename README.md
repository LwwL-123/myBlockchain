# myBlockchain
基于go语言实现的简单区块链demo

包含以下内容：
- hash
- Base58
- MerkleTree
- P2P
- POW
- boltDB
- UTXOS
- Transation
- Account

## p2p节点通信启动
``` go
// 1. 首先启动服务端监听
go run server.go

// 2. 启动客户端,依次为 姓名，ip，服务器端口号，本地端口号
go run client.go lzw localhost 9527 8081
go run client.go ww  localhost 9527 8082
