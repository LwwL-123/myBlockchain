package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 1.服务器启动监听
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 9527})
	defer listener.Close()
	fmt.Println("启动服务端在", listener.LocalAddr().String())

	// 定义切片存放两个udp地址
	peers := make([]*net.UDPAddr, 2, 2)
	buf := make([]byte, 256)

	// 2.从两个UDP消息中获得连接的地址A,B
	n, addr, _ := listener.ReadFromUDP(buf)
	fmt.Printf("read from < %s >: %s\n", addr.String(), buf[:n])
	peers[0] = addr

	n, addr, _ = listener.ReadFromUDP(buf)
	fmt.Printf("read from < %s >: %s\n", addr.String(), buf[:n])
	peers[1] = addr

	fmt.Println("begin nat \n")

	// 3. 将A和B分别介绍给彼此
	listener.WriteToUDP([]byte(peers[0].String()), peers[1])
	listener.WriteToUDP([]byte(peers[1].String()), peers[0])

	// 4.睡眠10s 确保发送成功，后关闭listener
	time.Sleep(time.Second * 10)
}
