package main

// make 关键字的主要作用是初始化内置的数据结构，也就是我们在前面提到的数组、切片和 Channel，而当我们想要获取指向某个类型的指针时可以使用 new 关键字，
// 只是知道如何使用 new 的人真的比较少，下面我们就来介绍一下 make 和 new 它们的区别以及实现原理

func main() {
	//虽然 make 和 new 都是能够用于初始化数据结构，但是它们两者能够初始化的结构类型却有着较大的不同，make 在Go语言中只能用于初始化语言中的基本类型：
	
	// slice := make([]int,0,100)
	// hash := make(map[int]bool,10)
	// ch := make(chan int,5)
	//
	
	
	
}
