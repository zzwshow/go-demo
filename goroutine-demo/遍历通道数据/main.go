package main

import (
	"fmt"
	"time"
)

func main() {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			// 向管道内发送3到0之间的数值
			ch <- i
			// 每次发完等待
			time.Sleep(time.Second)
		}
	}()

	// main 内遍历ch 数值
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)

		if data == 0 {
			break
		}
	}
}
