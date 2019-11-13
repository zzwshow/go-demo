package main

import (
	"fmt"
	"time"
)

//channel频率控制
//在对channel进行读写的时，go还提供了非常人性化的操作，那就是对读写的频率控制，通过time.Ticke实现




func main(){
	requests := make(chan int,10)
	for i :=0; i<10; i++ {
		requests <- i
	
	}
	close(requests)
	
	limitTime := time.Tick(time.Second*1)  // time.Tick 函数会返回 一个时间管道
	
	for v := range requests {
		<- limitTime                      // 加在循环体内 可以限制从管道内获取数据的频率
		fmt.Println("管道值：",v)
	}

}
