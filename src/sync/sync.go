package main

import (
	"fmt"
	"sync"
)

func main() {
	//var a = 0
	//for i := 0; i < 1000; i++ {
	//	go func(idx int) {
	//		a += 1
	//		fmt.Println(a)
	//	}(i)
	//}
	//time.Sleep(time.Second)

	// 互斥锁
	var a = 0
	var lock sync.Mutex
	for i := 0; i < 100; i++ {
		go func(idx int) {
			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Println(a)
		}(i)
	}

}
