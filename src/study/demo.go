package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	i2 := btoi(true)
	fmt.Println(i2)
	fmt.Println(btob(2))
	stra()
}

// bool 类型   true  false
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
func btob(i int) bool {
	a := 1
	if a > i {
		return true
	}
	return false
}

// 字符创
func stra() {
	var a = "123"
	fmt.Println(a)
	a = "456"
	fmt.Println(a)
	for e := range a {
		fmt.Println(e)
	}
	str := "你好啊alabama"
	fmt.Println(utf8.RuneCountInString(a))
	fmt.Println(utf8.RuneCountInString(str))

}
