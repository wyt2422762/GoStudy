package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务端")
	//监听8081端口
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("服务端启动失败")
		return
	}

	for {
		//接受客户端的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受客户端连接失败")
			return
		}

		go optConn(conn)
	}

}

func optConn(conn net.Conn){
	var buffer [256]byte
	for{
		_, err := conn.Read(buffer[:])
		if err != nil {
			fmt.Println("读取数据失败")
			return
		}
		fmt.Println("读取到数据，", string(buffer[:]))
	}
}
