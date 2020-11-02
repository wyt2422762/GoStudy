package main

import "fmt"

//切片(slice)
//Go中的切片是一个相同元素类型的可变长度序列
//Go中的切片是引用类型， 都指向了底层的数组

func main() {

	//声明一个int切片
	var s []int
	s = []int{1, 4} //初始化
	fmt.Println(s)
	fmt.Printf("len(s1): %d cap(s1): %d\n", len(s), cap(s))
	fmt.Println("---------------------")

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	s1 := a1[1:5] //得到数组的一个切片，左包含，右不包含 [2 3 4 5]
	s2 := a1[2:]  //得到数组的一个切片，左包含，右不包含 [3 4 5 6]
	s3 := a1[:1]  //得到数组的一个切片，左包含，右不包含 [1]
	s4 := a1[:]   //得到数组的一个切片，左包含，右不包含 [1 2 3 4 5 6]
	fmt.Println(s1, s2, s3, s4)
	//切片的容量cap指的是从切片的第一个元素到底层数组的最后一个元素的数量
	//切片的长度就是它元素的个数
	fmt.Printf("len(s1): %d cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s3): %d cap(s3): %d\n", len(s3), cap(s3))
	fmt.Printf("len(s4): %d cap(s4): %d\n", len(s4), cap(s4))
	fmt.Println("---------------------")

	//切片再切片
	s5 := s1[1:3]
	fmt.Println(s5)
	fmt.Printf("len(s5): %d cap(s5): %d\n", len(s5), cap(s5))
	fmt.Println("---------------------")

	//切片引用类型
	s5[0] = 100
	fmt.Println(a1, s1, s2, s3, s4, s5)
	fmt.Println("---------------------")

	//使用make函数构造切片 make(数据类型，切片长度，切片容量(不传该参数默认容量=长度))
	s6 := make([]int, 0, 10)
	s7 := make([]int, 5)
	fmt.Printf("len(s6): %d cap(s6): %d\n", len(s6), cap(s6))
	fmt.Printf("len(s7): %d cap(s7): %d\n", len(s7), cap(s7))
	fmt.Println(s6, s7)
	fmt.Println("---------------------")

	//切片扩容，append函数
	//必须用切片来接收append函数的返回值
	s8 := []string{"0", "1", "2", "3"}
	s8 = append(s8, "4")
	fmt.Println(s8)
	fmt.Println("---------------------")

	//切片拷贝
	s9 := []string{"0", "1", "2", "3"}
	var s10 = make([]string, len(s9))
	copy(s10, s9)
	fmt.Println(s9, s10)
	s9[1] = "100"
	fmt.Println(s9, s10)
	fmt.Println("---------------------")

	//切片删除
	s11 := []int{1, 2, 4, 5}
	s11 = append(s11[0:2], s11[3:]...) //...表示拆开
	fmt.Println(s11)
	fmt.Printf("len(s11): %d cap(s11): %d\n", len(s11), cap(s11))
	fmt.Println("---------------------")

}
