package main

import (
	"container/list"
	"fmt"
	"sort"
	"sync"
)

func main() {
	//	1.  map是一对元素（pair） 的无序集合

	//var mapname map[string]int // mapname 为map的变量名  [string] 为key 的类型  int 是value 类型
	//ints := make(map[string]int, 100)          可以执行map 的大小  例如100 		当map 容量达到上限时候，map的大小会会自动+1

	//
	//var map1 map[string]int
	//map1 = map[string]int{"one": 1, "two": 2} //
	//
	//map2 := make(map[string]int)
	//map2 = map1

	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	//2. 遍历 map  使用 for range
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
		if k == "route" {
			fmt.Println("111111")
		}
	}

	// 只遍历 value 时也可以使用 _ 方式遍历map
	for _, v := range scene {
		fmt.Println(v)
	}
	// 只遍历 key
	for k := range scene {
		fmt.Println(k)
	}

	//3. 对 map 中的值 进行排序
	var sceneList []string //  新建一个 切片
	for k := range scene {
		sceneList = append(sceneList, k) //将 key 放入list 中
	}

	sort.Strings(sceneList) // 对 切片 进行排序
	sort.Strings(sceneList)
	fmt.Println(sceneList)

	// 4.map 的删除 和清空
	// 使用 delete(map,键)的 方式

	scene2 := make(map[string]int, 20)
	scene2["one"] = 1
	scene2["two"] = 2
	scene2["third"] = 3
	scene2["four"] = 4
	for k := range scene2 {
		fmt.Println(k)
	}
	delete(scene2, "one")
	for k := range scene2 {
		fmt.Println(k)
	}

	//5 Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。
	//需要并发读写时，一般的做法是加锁，但这样性能并不高，Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map，sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。
	//
	//sync.Map 有以下特性：
	//无须初始化，直接声明即可。
	//sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
	//使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false

	var syncMap sync.Map
	// 添加几个键值
	syncMap.Store("one", 1)
	syncMap.Store("two", 2)
	syncMap.Store("third", 3)

	value, ok := syncMap.Load("one")
	fmt.Println(value, ok)
	fmt.Println(syncMap.Load("one"))
	syncMap.Delete("two")

	// 遍历 键值对
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	// 6. 在Go语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作。
	//list 初始化  var  变量名 = list.New()
	var list1 = list.New()
	//双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
	//这两个方法都会返回一个 *list.Element 结构，如果在以后的使用中需要删除插入的元素，则只能通过 *list.Element 配合 Remove() 方法进行删除，这种方法可以让删除更加效率化，同时也是双链表特性之一。
	list1.PushBack(1)
	list1.PushFront("dfadsfa")
	//list2 := list1.PushBack("first")
	// 遍历链表   遍历双链表需要配合 Front() 函数获取头元素，遍历时只要元素不为空就可以继续进行，每一次遍历都会调用元素的 Next() 函数，代码如下所示。
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
