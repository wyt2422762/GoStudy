package main

import "fmt"

//整型

func main() {
	var i = 10            //声明一个int类型(int类型的长度会根据操作系统的位数变化)
	fmt.Printf("%d\n", i) //10进制输出
	fmt.Printf("%o\n", i) //8进制输出
	fmt.Printf("%x\n", i) //16进制输出
	fmt.Printf("%T\n", i) //输出数据类型
	fmt.Println("-------------------------")
	var i8 = int8(9)       //声明一个int8类型
	fmt.Printf("%d\n", i8) //10进制输出
	fmt.Printf("%o\n", i8) //8进制输出
	fmt.Printf("%x\n", i8) //16进制输出
	fmt.Printf("%T\n", i8) //输出数据类型
}
