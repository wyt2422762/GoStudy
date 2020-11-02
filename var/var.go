package main

import "fmt"

//变量

//用 var 关键字来声明变量

//声明变量 var 变量名(Go语言中推荐使用驼峰式命名 idStr) 变量类型
var name string //全局变量
var age int     //全局变量
var isOk bool   //全局变量

//批量声明变量
var (
	name1   string      //全局变量
	ageCode int    = 18 //全局变量
	isOk1   bool        //全局变量
)

//s3 := "s3" //简短变量申明只能在函数中使用

func main() {
	name = "王彦廷"
	age = 18
	isOk = true

	fmt.Println("name: ", name)

	var temp string //局部变量，Go语言中局部变量声明后必须被使用，否则编译报错
	temp = "temp"   //局部变量赋值

	fmt.Println("temp: ", temp)

	var s1 string = "s1" //声明变量同时赋值

	fmt.Println("s1: ", s1)

	var auto = "auto" //类型推导，根据变量值判断变量类型
	var isAuto = true //类型推导，根据变量值判断变量类型

	fmt.Println("auto: ", auto)
	fmt.Println("isAuto: ", isAuto)

	s3 := "s3" //简短变量声明赋值，只能在函数中使用

	fmt.Println("s3: ", s3)

	x, y := foo() //批量声明变量
	fmt.Println("x: ", x)
	fmt.Println("x: ", y)

	x1, y1 := 1, 2 //批量声明变量
	fmt.Println("x1: ", x1)
	fmt.Println("y1: ", y1)

	x2, _ := foo() //_可以声明匿名变量，匿名变量不会分配内存空间
	fmt.Println("x2: ", x2)

	_, y2 := 1, 2 ///_可以声明匿名变量，匿名变量不会分配内存空间
	fmt.Println("y2: ", y2)

}

func foo() (int, string) {
	return 10, "字符串"
}
