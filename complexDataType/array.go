package main

import "fmt"

//数组
//Go数组的长度是数组类型的一部分
//Go中的数组为值类型
//[n]*T表示指针数组，*[n]T表示数组的指针

func main() {

	//声明一个长度为3的int数组
	var a1 [3]int
	a1 = [3]int{1, 2, 3} //初始化a1

	//声明一个长度为3的int数组,未初始化
	var a2 [4]int

	//根据实际内容推断数组长度
	var a3 = [...]int{1, 2, 3, 4, 5}

	var s1 [3]string
	s1 = [3]string{"王", "彦", "廷"}

	s2 := [6]string{0: "张三", 4: "李四"}   //利用索引初是化数组
	s3 := [...]string{1: "王五", 3: "李四"} //利用索引初是化数组

	fmt.Printf("a1: %T, a2: %T, a3: %T\n", a1, a2, a3) //Go数组的长度是数组类型的一部分
	fmt.Println(a1, a2, s1, s2, s3)                    //数组不初始化，默认值为零值(bool为false，数值为0，字符串为"")

	//数组遍历
	for i, s := range s2 { //for range
		fmt.Printf("索引：%d, 值：%v\n", i, s)
	}
	for _, s := range s2 { //for range
		fmt.Printf("值：%v\n", s)
	}

	for i := 0; i < len(s3); i++ { //普通for
		fmt.Printf("索引：%d, 值：%v\n", i, s3[i])
	}

	//多维数组
	var a32 [3][2]int //定义一个[3][2]的2维int数组
	//初始化
	a32 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
	}
	//遍历多维数组
	for _, outter := range a32 {
		fmt.Println(outter)
		for _, inner := range outter {
			fmt.Println(inner)
		}
	}

	//数组是值类型
	var q1 [3]int
	q1 = [3]int{3, 2, 1}
	q2 := q1
	q2[0] = 0
	fmt.Println(q1, q2)

}
