package main

import "fmt"

// 接收值判断,通道是否关闭

func send(ch chan int){
	for i:=0; i<10;i++{
		ch <- i //将i 的值传到通道里
	}
	close(ch)
	// 与打开文件不同,不关闭会产生僵尸进程,chan属于程序内部,会被go 的垃圾回收机制回收的,可以关闭,不关闭也没事儿
}


func main(){
	var ch1 = make(chan int,100)
	go send(ch1) //启动一个goroutine

	// 第一种方式:利用for从通道获取值
	//for {
	//	ret,ok := <- ch1 //通过取值时的:ok判断,当通道关闭后,ok=False
	//	if !ok{
	//		break
	//	}
	//	fmt.Println("ret:",ret)
	//}

	// 第二种方式 for range从通道取值,range 内部会自判断
	for ret := range ch1{
		fmt.Println(ret)
	}


}
