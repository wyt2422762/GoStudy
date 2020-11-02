package main

import (
	"fmt"
	"strings"
)

//字符串
//Go语言中的字符串编码为utf-8
//Go语言中的字符串必须用""(char型用'')
//Go语言中字符有两种类型， byte(ASCII码的一个子父，例如英文、数字等)、runne(utf-8的字符， 例如中文等)
//Go语言中字符串声明

func main() {
	var str = "字符串"
	fmt.Println("字符串：", str)
	fmt.Printf("%T\n", str) //输出数据类型
	fmt.Println("-------------------------")
	var chr = '字'
	fmt.Printf("字符：%c\n", chr)
	fmt.Printf("%T\n", chr) //输出数据类型
	fmt.Println("-------------------------")
	//多行字符串
	var s2 = `第1行
第2行
第3行`
	fmt.Println("多行字符串：", s2)
	fmt.Printf("%T\n", s2) //输出数据类型
	fmt.Println("-------------------------")

	//字符串操作
	fmt.Println("字符串s2长度：", len(s2)) //len()获取字符串长度
	fmt.Println("-------------------------")
	//字符串拼接
	ss1, ss2 := "ss1", "ss2"
	ss3 := ss1 + "-" + ss2
	fmt.Println("ss3: ", ss3)
	fmt.Println("-------------------------")
	ss4 := fmt.Sprintf("ss3: %s-%s", ss1, ss2) //fmt.sprintf
	fmt.Println("ss4: ", ss4)
	fmt.Println("-------------------------")
	ss5 := fmt.Sprint(ss1, ss2) //fmt.sprint
	fmt.Println("ss5: ", ss5)
	fmt.Println("-------------------------")
	ss6 := fmt.Sprintln(ss1, ss2) //fmt.sprintln
	fmt.Println("ss6: ", ss6)
	fmt.Println("-------------------------")

	//字符串分割
	var ori = "a|b|c"
	split := strings.Split(ori, "|")
	fmt.Println("字符串分割：", split)
	fmt.Println("-------------------------")

	//字符串宝航
	container := "这是一个测试一个字符串"
	con := strings.Contains(container, "测试")
	fmt.Println("字符串包含(测试)：", con)
	fmt.Println("-------------------------")
	con = strings.Contains(container, "测试1")
	fmt.Println("字符串包含(测试1)：", con)
	fmt.Println("-------------------------")

	//前缀、后缀
	con = strings.HasPrefix(container, "这是")
	fmt.Println("字符串前缀(这是)：", con)
	fmt.Println("-------------------------")
	con = strings.HasSuffix(container, "字符串")
	fmt.Println("字符串后缀(字符串)：", con)
	fmt.Println("-------------------------")

	//index lastIndex
	container = "abcbdefgab"
	indx := strings.Index(container, "b")
	fmt.Println("Index：", indx)
	fmt.Println("-------------------------")
	indx = strings.LastIndex(container, "a")
	fmt.Println("LastIndex：", indx)
	fmt.Println("-------------------------")

	//join
	fmt.Println("join with _ : ", strings.Join(split, "_"))
	fmt.Println("-------------------------")

	//字符串循环
	showStr := "abcs王彦廷"
	for i := 0; i < len(showStr); i++ {
		fmt.Printf("%v|%c\n ", showStr[i], showStr[i])
	}
	fmt.Println("-------------------------")
	for iii, v := range showStr {
		fmt.Printf("%d|%v|%c\n ", iii, v, v)
	}
	fmt.Println("-------------------------")

	ssss := "123123"
	ssss2 := []rune(ssss) //把字符串转成runne切片
	ssss2[0] = '0'
	fmt.Println(string(ssss2)) //把runne切片转成字符串

}
