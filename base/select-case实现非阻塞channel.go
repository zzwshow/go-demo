package main

import (
	"fmt"
	"time"
)

// select-case实现非阻塞channel

// 原理通过select+case加入一组管道，当满足（这里说的满足意思是有数据可读或者可写)select中的某个case时候，那么该case返回，若都不满足case，则走default分支。

func Send(c chan int){
	for i:=0;i<10;i++{
		c <- i
		fmt.Println("向通道内写入数据：",i)
	}

}


func main(){
	resCh := make(chan int, 10)
	strCh := make(chan string,10)
	go Send(resCh)    // 放入协程内跑
	time.Sleep(time.Second)  // 加这个是为了 select 可以匹配到resCh ，如果不加主程序内的 strCh <- "hello" 先执行select 就会匹配到b:= <-strCh
	strCh <- "hello"
	
	select {
	case a:= <-resCh:
		fmt.Println("get data:",a)
	case b:= <-strCh:
		fmt.Println("Get data:",b)
	default:
		fmt.Println("默认。。。。")
		
	}

}
