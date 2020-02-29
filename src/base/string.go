package main

func main() {
	// 		\n：换行符
	//		\r：回车符
	//		\t：tab 键
	//		\u 或 \U：Unicode 字符
	//		\\：反斜杠自身
	var a = "111111111"
	var b = "dfd"
	println(a)

	//一般的比较运算符（==、!=、<、<=、>=、>）是通过在内存中按字节比较来实现字符串比较的，因此比较的结果是字符串自然编码的顺序。字符串所占的字节长度可以通过函数 len() 来获取，例如 len(str)。

	println(a == b, len(a), len(b))
	println(a > b)

	//	在Go语言中，使用双引号书写字符串的方式是字符串常见表达方式之一，被称为字符串字面量（string literal），这种双引号字面量不能跨行，如果想要在源码中嵌入一个多行字符串时，就必须使用`反引号
	//在`间的所有代码均不会被编译器识别，而只是作为字符串的一部分。

	const str = `第一行
第二行
第三行`
	println(str)
}