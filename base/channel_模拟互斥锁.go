package main

import (
	"fmt"
	"sync"
	"time"
)

//// 使用channel 完成go程数据同步

//var ch = make(chan bool)
//
//func printer(str string){
//	for _,s := range str {
//		fmt.Printf("%c",s)
//		time.Sleep(time.Microsecond* 300)
//	}
//}
//
//
//func person1(){
//	printer("hello")
//	ch <- true   			// 1 使用通道解决, 当person1 执行完printer 函数后向通道内写入数据
//	close(ch)
//}
//
//func person2(){
//	<- ch        // 2 若收不到数据,则阻塞当前go子程,下边代码不在执行,若读取到数据,则继续执行
//	printer("world")
//
//}
//
//func main(){
//	go person1()   // 这两个子go程谁枪占到CPU 谁打印,最终打印的字符 //hweolrllod
//	go person2()
//	for  {
//		;
//	}
//
//}


//// 使用"锁" 完成同步
/*
	1,互斥锁
	2,读写锁


*/

var mutex sync.Mutex  // 创建一个互斥量(也就是互斥锁). 新建的互斥锁为0, 是未枷锁状态

func printer(str string){     // 由于两个子go程 的共享数据就是这个函数,所以在这个函数内枷锁,谁先来就或取到锁,然后执行,待函数执行完,在解锁
	mutex.Lock()              // 函数进来以后就枷锁
	for _,s := range str {
		fmt.Printf("%c",s)
		time.Sleep(time.Microsecond* 300)
	}
	mutex.Unlock()				// 函数执行完之后 解锁
}

func person1(){
	printer("hello")
}

func person2(){
	printer("world")

}

func main(){
	go person1()   // 这两个子go程谁枪占到CPU 谁打印,最终打印的字符 //hweolrllod
	go person2()
	for  {
		;
	}

}