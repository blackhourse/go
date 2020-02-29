package main

import "fmt"

func main() {
	ddd()
}
func ddd() {
	//bool 只有 true false 两个值
	fmt.Println((!true == false) == true)
	//	&&（and ） || (or) 如果运算值左边的值已经可以满足整个布尔表达式的值， 那么运算符右边的值将不再被求值
	// && 的优先级比|| 的优先级高  （数学中 乘法的优先级 > 加法的优先级）

	orFalse := isTrueOrFalse(3)
	fmt.Println(orFalse)

}

// 根据值大小 返回 1 0
func isTrueOrFalse(i int) int {
	if i > 1 {
		return 1
	}
	return 0

}
