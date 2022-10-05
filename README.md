
# 实现简易的 netcat 的TCP 客户端


$ go run ./main.go www.baidu.com 80  # 连接到 www.baidu.com 80 端口，发送数据
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
	


flag 解析命令行



超时控制
    建立连接
        net.DialTimeout()

    读写数据
    

$ nc -v -w 3 www.google.cpm 80



flag 解析命令行

