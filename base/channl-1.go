package main

import "fmt"

// 无缓冲通道和缓冲通道 (缓冲通道,有容量值)
func recv(ch chan bool){
	ret := <-ch
	fmt.Println(ret)
}


func main() {
	ch := make(chan bool) //创建一个通道
	go recv(ch) // 起一个goroutine
	ch <- true //往通道内传值

}