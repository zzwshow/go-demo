package goroutine_01

import (
	"fmt"
	"sync"
)

// goroutine
var wg sync.WaitGroup //优雅的等待goroutine执行完毕!

func hello(){
	//防止函数出错,先注册一个defer,防止wg.Done() 不执行,防止死锁!
	defer wg.Done() //执行完计数器减一,
	fmt.Println("今天好冷~~~~")
}

func main(){

	wg.Add(5) //增加计数器,就是启动了多少个goroutine
	for i:=0;i<5;i++{
		go hello() //启动5个goroutine
	}
	//使用sleep 等待
	//time.Sleep(time.Second) //会让cpu 真正的停止,尽量少用!

	fmt.Println("今天吃什么")
	wg.Wait() //等待所有goroutine 等执行完
}
