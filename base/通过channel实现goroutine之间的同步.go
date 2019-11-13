package main

import (
	"fmt"
	"time"
)

//通过channel实现goroutine之间的同步。
/*
实现方式：
	通过channel能在多个groutine之间通讯，当一个goroutine完成时候向channel发送退出信号,等所有goroutine退出时候，
利用for循环channe去channel中的信号，若取不到数据会阻塞原理，等待所有goroutine执行完毕，
使用该方法有个前提是你已经知道了你启动了多少个goroutine。
*/

func cal(a, b int ,Exitchan chan bool) { //参数为两个int 整型，和一个接收bool 类型的通道
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
	time.Sleep(time.Second*2)
	Exitchan <- true  // 向通道内传值
}


func main(){
	Exitchan := make(chan bool, 10) // 声明并分配管道内存，使用缓存通道，缓存10个元素
	
	for i:=0; i<10; i++{
		go cal(i,i+1,Exitchan)  //创建10个groutine, （这个执行后会向管道内传入10个true元素）
	}
	
	
	// 定义一个接收管道元素的函数
	for j:=0;j<10;j++{
		// <- Exitchan  //取通道内的数据元素，如果取不到则会阻塞（这里取出后直接丢弃了）
		tmp := <- Exitchan
		fmt.Printf("通道内的第%d 个元素：%v\n", j, tmp)
	}
	close(Exitchan) // 关闭通道

}
