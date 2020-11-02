package main

import "fmt"

//for语句
//Go语言中只有for一种循环

func main() {

	//基本for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	//变种1
	i1 := 2
	for ; i1 < 10; i1++ {
		fmt.Printf("i1 = %d\n", i1)
	}

	//变种2
	i2 := 5
	for i2 < 10 {
		fmt.Printf("i2 = %d\n", i2)
		i2++
	}

	//死循环
	/*for {
		//do something
	}*/

	//for range i 索引(key) v(value)
	str := "Hello,王彦廷"
	for i, v := range str {
		fmt.Printf("i = %d, v = %c\n", i, v)
	}

	//使用goto|break + label 实现跳出多层循环
	//cusLabel1:
	for i3 := 0; i3 < 10; i3++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d - %d\n", i3, j)
			if j == 4 {
				goto cusLabel2 //跳到指定标签位置后向下执行(尽量不要使用)
				//break cusLabel1 //利用标签跳出多层循环
			}
		}
	}
cusLabel2: //定义一个标签
}
