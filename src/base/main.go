package main

import (
	"fmt"
)

var i int = 100

func main() {
	//fmt.Print("hello,world,go-base")
	//var a int = 100
	//var b int = 200
	//b, a = a, b
	//fmt.Println(a, b)
	_, b := GetData()
	_, a := GetData()
	fmt.Println(a, b)
	Add()
	i = 1200
	fmt.Println("i:", i)
}

//func (p IntSlice) Len() int           { return len(p) }

func GetData() (int, int) {
	return 100, 200
}
func Add() {
	var a int = 2
	var b int = 2
	c := a + b
	fmt.Printf("a = %d , b = %d ,c = %d ", a, b, c)
}
