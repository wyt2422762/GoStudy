package main

import "fmt"

//常量计数器

//iota是Go语言的常量计数器，只能在常量表达式中使用
//iota在const关键字出现是被置为0，const中每增加一行常量声明iota增加1(iota可以理解为const块中的行索引)

const iotaN1 = iota //0
const iotaN2 = iota //0

const (
	iotaD1 = iota //iota = 0
	iotaD2 = iota //iota = 1
	iotaD3 = iota //iota = 2
)

const (
	iotaE1 = iota //iota = 0
	iotaE2 = 100  //iota = 1
	_      = iota //iota = 2
	iotaE3 = iota //iota = 3
	iotaE4        //iota = 4
)

const (
	iotaF1, iotaF2 = iota, iota //iota = 0
	iotaF3, iotaF4 = iota, iota //iota = 1
)

//iota用法示例1：定义数量级
const (
	_  = iota             //iota = 0
	KB = 1 << (10 * iota) //iota = 1 2的10次方
	MB = 1 << (10 * iota) //iota = 2 2的20次方
	GB = 1 << (10 * iota) //iota = 3 2的30次方
	TB = 1 << (10 * iota) //iota = 4 2的40次方
	PB = 1 << (10 * iota) //iota = 5 2的50次方
)

func main() {
	fmt.Println("iotaN1:", iotaN1)
	fmt.Println("iotaN2:", iotaN2)
	fmt.Println("---------------")
	fmt.Println("iotaD1:", iotaD1)
	fmt.Println("iotaD2:", iotaD2)
	fmt.Println("iotaD3:", iotaD3)
	fmt.Println("---------------")
	fmt.Println("iotaE1:", iotaE1)
	fmt.Println("iotaE2:", iotaE2)
	fmt.Println("iotaE3:", iotaE3)
	fmt.Println("iotaE4:", iotaE4)
	fmt.Println("---------------")
	fmt.Println("iotaF1:", iotaF1)
	fmt.Println("iotaF2:", iotaF2)
	fmt.Println("iotaF3:", iotaF3)
	fmt.Println("iotaF4:", iotaF4)
	fmt.Println("---------------")
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
}
