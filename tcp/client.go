package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端")

	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		fmt.Println("与服务端简历连接失败")
	}

	for {
		fmt.Println("请输入：")
		reader := bufio.NewReader(os.Stdin)
		byt, _, _ := reader.ReadLine()

		_, err := conn.Write(byt)
		if err != nil {
			fmt.Println("向服务端写数据失败")
		}
	}

}
