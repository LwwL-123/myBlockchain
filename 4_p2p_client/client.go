package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1.设置参数
	if len(os.Args) < 5 {
		fmt.Println("./Client tag remoteIP remotePort port")
		return
	}
	// 客户端标识
	tag := os.Args[1]
	// 服务器IP
	remoteIP := os.Args[2]
	// 服务器端口
	remotePort, _ := strconv.Atoi(os.Args[3])
	// 本地绑定端口   strconv.Atoi 字符串转整型
	port, _ := strconv.Atoi(os.Args[4])
	// 绑定本地端口
	localAddr := net.UDPAddr{Port: port}

	// 2. 与服务器建立连接，发消息
	conn, err := net.DialUDP("udp", &localAddr, &net.UDPAddr{IP: net.ParseIP(remoteIP), Port: remotePort})
	if err != nil {
		log.Panic("UDP连接失败", err)
	}
	// 2.1 发送消息，自我介绍
	conn.Write([]byte("我是" + tag))

	// 3.从服务器获得目标地址
	buf := make([]byte, 256)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Panic("读取udp连接失败", err)
	}
	conn.Close()
	toAddr := parseAddr(string(buf[:n]))
	fmt.Println("获得连接对象地址:", toAddr)

	// 4.建立p2p连接
	p2p(&localAddr,&toAddr)
}

// 实现p2p通信
func p2p(srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) {
	// 1. 请求建立连接
	conn, _ := net.DialUDP("udp", srcAddr, dstAddr)
	// 2. 发送消息
	conn.Write([]byte("打洞消息\n"))

	// 3.启动一个goroutine监控输入标准
	go func() {
		buf := make([]byte, 256)
		for {
			// 接收udp消息
			n, _, _ := conn.ReadFromUDP(buf)
			if n > 0 {
				fmt.Printf("接收消息: %sp2p > ",buf[:n])
			}
		}
	}()

	// 4.监控输入标准，发送给对方
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("p2p > ")
		// 读取标准输入，已换行为读取标志
		data, _ := reader.ReadString('\n')
		conn.Write([]byte(data))
	}
}

// 解析地址函数
func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}
