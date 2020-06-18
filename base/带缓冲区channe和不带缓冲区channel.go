package main

import (
	"fmt"
	"time"
)

//带缓冲区channe和不带缓冲区channel

/*
带缓冲区channel：定义声明的时候指定了缓冲区大小(长度)，可以保存多个数据。
不带缓冲区channel：只能存一个数据，并且只有当该数据被取出时候才能存下一个数据。

	ch := make(chan int)  //不带缓冲区
	ch1 := make(chan int,10)  // 带缓冲区
*/


func send(s chan int) {
	time.Sleep(time.Second*3)
	for i:=0;i<10;i++{
		fmt.Println("send data:",i)
		s <-i
	}
	close(s)  // 因为在主程序内使用range  循环读取管道，若管道不关闭，range读完后或产生死锁
}

func main(){
	sChan := make(chan int)
	go send(sChan)
	
	for v := range sChan{
		fmt.Printf("接收通道数据：%d\n",v)
	}


	

}

