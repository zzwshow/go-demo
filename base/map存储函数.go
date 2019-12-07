package main

import (
	"runtime"
	"sync"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

//定义函数类型

type Msg func(name string) string

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // 此程序使用的最大逻辑CPU 数量
	wg := &sync.WaitGroup{}
	c := make(chan os.Signal, 1)
	handleMap := make(map[int]Msg)
	handleMap[1] = handle1
	handleMap[2] = handle2
	handleMap[3] = handle3
	fmt.Println("执行任务～～～")
	
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)  // 在c管道内 传递系统输入信号
	go func() {
		s := handleMap[3]
		data := s("测试")
		fmt.Println(data)
		<-c       // 接收信号
		wg.Done()  // goroutine 执行完成之后 计数器减一
		
	}()
	
	wg.Add(1)   // 计数器
	
	wg.Wait()         // 等待任务都执行完在退出, 在有goroutine 没有执行完之前 阻塞主线程
	fmt.Printf("结束")
}

func handle1(name string) string {
	fmt.Println(name)
	return "handle1"
	
}
func handle2(name string) string {
	fmt.Println("handle2")
	return "handle2"
	
}
func handle3(tt string) string {
	fmt.Println(tt)
	return "handle3"
	
}
