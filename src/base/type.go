package main

import "fmt"

type NewInt int     //将 NewInt 定义为 int 类型，这是常见的定义类型的方法，通过 type 关键字的定义，NewInt 会形成一种新的类型，NewInt 本身依然具备 int 类型的特性。
type IntAlias = int //将 IntAlias 设置为 int 的一个别名，使用 IntAlias 与 int 等效。

func main() {

	//区分类型别名与类型定义
	var a NewInt // 将a 声明为 NewInt 类型
	fmt.Printf("a type:%T\n", a)
	var a2 IntAlias
	fmt.Printf("a2 type:%T\n", a2)
	fmt.Println(a)
	fmt.Println(a2)
}
