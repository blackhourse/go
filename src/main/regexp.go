package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	//regexp1()
	test()
}

func regexp1() {
	buf := "abc azc a7c aac 888 a9c  tac"
	compile := regexp.MustCompile("a.c")
	if compile == nil {
		fmt.Println("匹配空")
	}
	submatch := compile.FindAllStringSubmatch(buf, -1)
	fmt.Println(submatch)
}

func test() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为 "##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
