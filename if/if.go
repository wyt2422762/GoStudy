package main

import "fmt"

//if语句

func main() {

	age := 19

	//基本if语句
	if age >= 18 {
		fmt.Println("成年")
	} else if age > 10 {
		fmt.Println("少年")
	} else {
		fmt.Println("小孩")
	}

	//if特殊写法, age1只在if的作用域内生效
	if age1 := 20; age1 >= 18 {
		fmt.Println("成年")
	} else if age > 10 {
		fmt.Println("少年")
	} else {
		fmt.Println("小孩")
	}
	//if作用域外无age1对象
	//fmt.Println(age1)

}
