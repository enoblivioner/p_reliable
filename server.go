package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// 处理错误信息
func checkError(err error) {
	if err != nil { //指针不为空
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
}

// 单次接收UDP报文，并在接收缓冲区中确认
func receiveUDPMsg(udpConn *net.UDPConn) {

	//声明单次接收缓冲区
	buffer := make([]byte, 30)

	//从udpConn读取客户端发送过来的数据，放在缓冲区中(阻塞方法)
	//返回值：n=读到的字节长度,remoteAddr=客户端地址,err=错误
	n, remoteAddr, err := udpConn.ReadFromUDP(buffer) //从udp接收数据读取到buffer中
	checkError(err)

	//接收方运行时查看实时报文接收情况
	fmt.Printf("接收到来自%v的消息:%s,大小为%d字节\n", remoteAddr, string(buffer[0:n]), n)

	if string(buffer[0:n]) == "clientover" {
		udpConn.WriteToUDP([]byte("serverover"), remoteAddr)
		time.Sleep(time.Second)
		os.Exit(1)
	}
}

// 主函数
func main() {

	// 定义命令行参数
	prot := flag.String("prot", "udp", "protocol name")
	addr := flag.String("addr", "127.0.0.1:8848", "ip:port")
	flag.Parse()

	//解析IP和端口得到UDP地址
	udp_addr, err := net.ResolveUDPAddr(*prot, *addr)
	checkError(err)
	fmt.Println(udp_addr)

	//在解析得到的地址上建立UDP监听
	udpConn, err := net.ListenUDP("udp", udp_addr)
	checkError(err)
	fmt.Printf("在%s建立监听:···\n", *addr)
	defer udpConn.Close() //关闭连接
	
	//从udpConn中接收UDP消息，返回SACK
	for {
		receiveUDPMsg(udpConn)
	}

}
