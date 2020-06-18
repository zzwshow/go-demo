package main

import (
	"fmt"
	"sync"
)

//Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。
//sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。

// sync.Map 有以下特性：
// 无须初始化，直接声明即可。
// sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
// 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。


func main() {
	var scene sync.Map
	// 将键值对保存到sync.map
	
	scene.Store("greece",97)
	scene.Store("london",100)
	scene.Store("egypt",200)
	
	// 从sync.map 中根据建取值
	fmt.Println(scene.Load("greece"))
	
	// 根据键删除对应的键值对
	scene.Delete("egypt")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k,v interface{}) bool {
		fmt.Println("iterate:",k, v)
		return true
	})
	//Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回。
	
}
