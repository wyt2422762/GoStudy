package main

import "fmt"

//map

func main() {
	//声明一个key为int，value为string的map
	var m1 map[int]string //map刚声明时为nil，需要初始化后使用
	fmt.Println(m1 == nil)
	m1 = make(map[int]string, 10) //初始化map，分配内存空间
	m1[1] = "1"
	m1[2] = "2"
	fmt.Println(m1)
	fmt.Println("---------------------")

	//判断key是否存在
	v, ok := m1[3]
	if !ok {
		fmt.Printf("key存在, 值：%v\n", v)
		fmt.Println(v == "")
	} else {
		fmt.Printf("key存在, 值：%v\n", v)
	}
	fmt.Println("---------------------")

}
