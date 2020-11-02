package main

import "fmt"

//指针
//&获取地址 *根据地址获取值

//Go语言中 new和make都可以分配内存空间
//new 和 make 的区别
//1. new很少用，一般用来给基本数据类型申请内存，返回类型为对应类型的指针 new(int) -> *int
//2. make是专门用来给切片(slice)，map，chan类型申请内存的，返回类型为对应的三个类型自身
//3. 二者都是用来做内存分配的

func main() {

	n1 := 1
	p := &n1
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	n2 := *p
	fmt.Println(n2)
	fmt.Println(n2 == n1)
	fmt.Printf("%T\n", n2)
	fmt.Println("---------------------")

	/*var a *int //声明一个int指针，还未初始化，次是值为nil
	fmt.Println(a == nil) //true
	*a = 100 //该行会报空指针异常
	fmt.Println(*a)*/

	var a1 = new(int)      //new函数分配一块内存地址
	fmt.Println(a1 == nil) //false
	fmt.Printf("%T\n", a1)
	fmt.Println(a1)
	fmt.Println(*a1)
	*a1 = 100
	fmt.Println(a1)
	fmt.Println(*a1)
	fmt.Println("---------------------")

}
