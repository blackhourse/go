package main

import (
	"fmt"
	"strconv"
)

type Status string //   Status 是一个自定义枚举类  string 是枚举返回的类型

// 模拟枚举类型
const (
	SUCCESS Status = "success" // 开始生成枚举值, 默认为0
	FAIL           = "fail"
	EQUALS         = "equals"
)
const (
	//q = iota // iota 默认值是 0   并且 下个枚举 自动加1
	q = 1 << iota // 每个枚举想左进行位运算2位
	b
	c
	d
	f
	g
)

func main() {

	// 输出所有枚举值
	fmt.Println(SUCCESS, FAIL, EQUALS)
	// 使用枚举类型并赋初值
	var a Status = EQUALS
	fmt.Println(a)
	fmt.Println(q, b, c, d, f, g)

	// 整型 转字符串   Itoa()：整型转字符串
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("type:%T\n", num)
	fmt.Printf("type:%T", str)

	//	Atoi() 字符串转整形
	str1 := "100"
	str2 := "s100"
	num1, error := strconv.Atoi(str1)
	if error != nil {
		fmt.Printf("%v 转换失败", num1)
	} else {
		fmt.Printf("type:%T value:%#v\n", num1, num1)
	}
	num2, error := strconv.Atoi(str2)
	if error != nil {
		fmt.Printf("%v 转换失败", num2)
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)

	}

}
