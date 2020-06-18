package main

import (
	"fmt"
	"runtime"
	"time"
)

// 通道超时控制




func main(){
	ch := make(chan int)
	quit := make(chan bool)

	go func() {  // 子go程获取数据
		for {
			select {
			case data := <-ch: //select 监听管道数据流  如果有数据就打印
				fmt.Println("data = ",data)
			case <-time.After(time.Second *3):  // 若3秒内,没有从管道中获取数据,就向 quit管道发送bool数据 通知主线程退出
				quit <- true
				runtime.Goexit() //可以直接跳出for 循环
				// break //只能跳出select
				// return // 可以直接跳出for 循环

			}
		}

	}()

	for i:=0; i<2; i++{  //向ch 管道发送数据
		ch <- i
		time.Sleep(time.Second * 2)		// 不能超过三秒,超过三秒select会匹配到第二个case,退出主go程
	}

	<-quit  //主go程 接收quit管道数据,若quit内没有数据,则阻塞主go程,收到则获取后丢弃并退出主go程
	fmt.Println("Finish!!")

}
