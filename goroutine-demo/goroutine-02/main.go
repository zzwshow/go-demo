package main

import (
	"fmt"
	"runtime"
	"sync"
)


var wg sync.WaitGroup

func a(){
	defer wg.Done()
	for i:=0;i<10;i++{
		fmt.Println("-------A--------")
	}
}

func b(){
	defer wg.Done()
	for i:=0;i<10;i++{
		fmt.Println("-------B--------")
	}
	for{
		fmt.Println("lllllllllllllllllllll")
	}
}


func main() {
	runtime.GOMAXPROCS(1) //设置go 程序使用几个CPU逻辑核心
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
