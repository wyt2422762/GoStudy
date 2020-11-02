package main

import "fmt"

//布尔型
//Go语言中bool型默认值为false
//Go语言中不允许将int转成bool
//Go语言中bool无法参加数值运算，也无法与其他类型进行转换

func main() {
	var bl1 bool //定义一个bool型
	fmt.Println("bool1:", bl1)
	fmt.Println("----------------------------")
	bl2 := true
	fmt.Printf("%T : %v\n", bl2, bl2)

	//bl2 = (bool)1 //int强转bool会报错
	//fmt.Println("bool1:", bl2)

}
