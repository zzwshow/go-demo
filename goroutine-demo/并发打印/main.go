package main

import "fmt"

func printer(c chan int) {
	for {
		data := <-c
		fmt.Println(data)

		if data == 0 {
			break
		}
	}

	c <- 0

}

func main() {
	// 创建一个管道
	c := make(chan int)

	// 并发执行 printer 传入channel
	go printer(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}

	c <- 0

	<-c

}
