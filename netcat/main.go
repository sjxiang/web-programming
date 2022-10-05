package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

/*
实现简易的 netcat TCP 客户都


nc www.baidu.com 80  # 连接到 www.baidu.com 80 端口，发送数据
GET / HTTP/1.1
\r\n

读取命令行参数 
	os.Args
	[]string{"nc", "www.baidu.com", "80"}


建立 TCP 连接
	net.Dial()


网络数据读写
	net.Conn
		Read 读取
		Write 写入
		Close 回收 IO

	io.Copy 双工 阻塞
		
		*/



func main() {
	fmt.Println(os.Args)

	if len(os.Args) != 3 {
		log.Fatal("用法：nc host port")
	}

	host := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", host + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()  // 发送 FIN 包，指示目标端口断开连接

	go func ()  {
		// 读取标准输入数据，发送给目标端口
		io.Copy(conn, os.Stdin)	
	}()
	
	// 读取目标端口的数据，并写入到标准输出
	io.Copy(os.Stdout, conn)
}