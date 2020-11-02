package main

import "fmt"

//复数 complex128(默认) complex64

func main() {
	var complex1 = 1 + 2i //定义一个复数
	fmt.Println(complex1)
	fmt.Printf("%T\n", complex1) //输出数据类型
	fmt.Println("---------------------")
	var complex2 = complex64(2 + 4i) //定义一个complex64类型的复数
	fmt.Println(complex2)
	fmt.Printf("%T\n", complex2) //输出数据类型
}
