package main

import (
	"fmt"
)

func main() {

	// nil   不能用于比较      指针 切片 映射 函数  接口的 零值是 nil     nil 没有默认类型
	//fmt.Println(nil == nil) //.\nil.go:8:18: invalid operation: nil == nil (operator == not defined on nil)
	// nil 不是关键字 或保留字
	//nil := errors.New("test")  // 所以可以定义为变量名称
	//fmt.Println(nil)

	// 不同类型nil 的指针是一样的
	var arr []int
	var num *int
	fmt.Println("%p\n", arr)
	fmt.Println("%p", num)

	//	不同类型的 nil 值占用的内存大小可能是不一样的
	//nil 是 map、slice、pointer、channel、func、interface 的零值
}
