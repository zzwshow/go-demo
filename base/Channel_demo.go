package main

import "fmt"

//channel

/*
channel俗称管道，用于数据传递或数据共享，其本质是一个先进先出的队列，使用goroutine+channel进行数据通讯简单高效，同时也线程安全，
多个goroutine可同时修改一个channel，不需要加锁。

channel可分为三种类型：
	只读channel：只能读channel里面数据，不可写入
	只写channel：只能写数据，不可读
	一般channel：可读可写
*/


/*
定义和声明
	var readOnlyChan <-chan int  // 只读通道
	var writOnlyChan chan<- int  // 只写通道
	var myChan chan int          // 可读可写通道

已上定义完不能直接使用，因为还没有分配内存空间，需要用make来分配，直接使用会造成deadlock 死锁
	read_only := make (<-chan int,10)//定义只读的channel
	write_only := make (chan<- int,10)//定义只写的channel
	read_write := make (chan int,10)//可同时读写

操作
	ch <- "abc"  //写入管道
	tmp :=  <- ch   //读取管道数据，并复值给变量tmp
	tmp, ok := <- ch    //优雅的读取数据

读写数据 注意事项：
	1、管道如果未关闭，在读取超时会则会引发deadlock异常
	2、管道如果关闭进行写入数据会pannic
	3、当管道中没有数据时候再行读取或读取到默认值，如int类型默认值是0
 */


// 循环管道  需要注意的是：
// 1、使用range循环管道，如果管道未关闭会引发deadlock错误。
// 2、如果采用for死循环已经关闭的管道，当管道没有数据时候，读取的数据会是管道的默认值，并且循环不会退出。


func main(){
	myChan := make(chan int,100) // 定义管道（同时分配内存）
	for i:=0; i<10; i++{
		myChan <- i  //向管道内传值
	}
	close(myChan)  // 关闭管道  //上面的for 循环已经写入了10个数据到管道内，长度是：10
	fmt.Println("data length: ",len(myChan))  //打印管道长度
	for v := range myChan{      // 使用range 循环从管道内取值   , 此时管道已经是关闭的，若没有关闭range 会产生死锁
		fmt.Println("V: ",v)
	}
	fmt.Printf("data length: %d",len(myChan))   // range 循环已经取出了管道内数据，故长度为 0



}













