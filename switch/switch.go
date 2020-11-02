package main

import "fmt"

//switch语句
//Go语言中switch中不需要加break，如果要下穿需要加 fallthrough

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		//fallthrough //表示hi执行下一个case语句
	case 3:
		fmt.Println(3)
	default:
	}

	switch i {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
		break
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
		break
	default:
		break
	}

	switch {
	case i == 1:
		fmt.Println("i == 1")
	case i == 2:
		fmt.Println("i == 2")
	case i == 3:
		fmt.Println("i == 2")
	}

}
