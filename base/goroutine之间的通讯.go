package main

import (
	"fmt"
	"sync"
)

//goroutine之间的通讯
/*
goroutine本质上是协程，可以理解为不受内核调度，而受go调度器管理的线程。goroutine之间可以通过channel进行通信或者说是数据共享，当然你也可以使用全局变量来进行数据共享。
*/
// 使用channel模拟消费者和生产者模式


func productor(myChan chan int, data int,wait *sync.WaitGroup) {
	myChan <- data
	fmt.Println("productor data: ",data)
	defer wait.Done()
}

func Consumer(myChan chan int, wait *sync.WaitGroup) {
	tmp := <- myChan
	fmt.Println("Consumer data: ",tmp)
	defer wait.Done()
}

func main(){
	
	myChan := make(chan int,100) // 创建管道
	var wg sync.WaitGroup  // 创建计数器对象
	
	for i:=0; i<10; i++{
		go productor(myChan,i,&wg)  //创建生产者 goroutine
		wg.Add(1)
	}
	
	for j:=0; j<10; j++{
		go Consumer(myChan,&wg)   // 创建消费者 goroutine
		wg.Add(1)
	}
	wg.Wait()
	
	close(myChan)

}
