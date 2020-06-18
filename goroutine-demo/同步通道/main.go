package main

import "fmt"

func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main 的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")
	// 等待匿名 goroutine
	<-ch // 阻塞main 直到从ch管道内接受到数据
	fmt.Println("call done")

}
