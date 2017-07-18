package main

import (
	"fmt"
	"net"
	"os"
)

func sender(conn net.Conn) {

	words := "hello world!"
	conn.Write([]byte(words))    //向tcpconn中写入数据
	fmt.Println("send over")

}

func main() {
	server := "127.0.0.1:1024"
	//net参数是"tcp4"、"tcp6"、"tcp"
	//addr表示域名或IP地址加端口号
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server) //获取一个TCPAddr
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	//net参数是"tcp4"、"tcp6"、"tcp"
	//laddr表示本机地址，一般设为nil
	//raddr表示远程地址
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //建立一个TCP连接
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

}
