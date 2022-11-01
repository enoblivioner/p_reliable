package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)



func main() {
	//命令行参数
	prot := flag.String("prot", "udp", "protocol name")
	addr := flag.String("addr", "127.0.0.1:8848", "ip:port")
	cwnd := 100
	flag.Parse()

	//请求连接服务器，得到连接对象
	conn, err := net.Dial(*prot, *addr)
	if err != nil {
		fmt.Println("网络连接出错")
		os.Exit(1)
	}

	defer conn.Close()

	//多次向连接中写入序号不同的报文
	string1 := "号报文"
	for num := 0; num < cwnd; num++ {
		string2 := fmt.Sprintf("%d", num)
		string3 := string2 + string1
		conn.Write([]byte(string2))
		fmt.Println("发送消息", string3)
	}

	conn.Write([]byte("clientover"))

	//读取代表收取消息(阻塞)
	buffer := make([]byte, 30)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取消息错误: err=", err)
		os.Exit(1)
	}

	fmt.Println(string(buffer[0:n]))
	if string(buffer[0:n]) == "serverover" {
		fmt.Println("yes")
		os.Exit(1)
	}

	
}
