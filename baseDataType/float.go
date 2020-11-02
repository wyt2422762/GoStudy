package main

import "fmt"

//浮点型 float64 float32

func main() {
	var f1 = 3.1415926
	fmt.Printf("%f\n", f1)   //以浮点型输出
	fmt.Printf("%.2f\n", f1) //以浮点型输出，2位小数点
	fmt.Printf("%T\n", f1)   //输出数据类型
	fmt.Println("-------------------------")
	var f2 = float32(3.1415926) //定义一个float32类型的变量
	fmt.Printf("%f\n", f2)      //以浮点型输出
	fmt.Printf("%.9f\n", f1)    //以浮点型输出，9位小数点
	fmt.Printf("%T\n", f2)      //输出数据类型

}
