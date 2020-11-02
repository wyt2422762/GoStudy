package main

import "fmt"

//常量

const pi = 3.1415926 //常量声明的时候必须赋值，且赋值后值不会也不能发生变化
//pi = 123 //常量不能重新赋值

//批量声明常量
const (
	constN1 = 100
	constN2 = 50
	constN3 = 25
)

//声明多个变量时，省略赋值操作则表示与上一行的常量值相同(因此首行的常量必须有赋值操作)
const (
	//d0
	constD1 = 25
	constD2
	constD3 = 30
	constD4
	constD5
)

func main() {
	//pi = 123123 //常量不能重新赋值
	fmt.Println("pi:", pi)

	fmt.Println("constN1:", constN1)
	fmt.Println("constN2:", constN2)
	fmt.Println("constN3:", constN3)

	fmt.Println("constD1:", constD1)
	fmt.Println("constD2:", constD2)
	fmt.Println("constD3:", constD3)
	fmt.Println("constD4:", constD4)
	fmt.Println("constD5:", constD5)

}
