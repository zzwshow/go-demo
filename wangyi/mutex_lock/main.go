package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var wg sync.WaitGroup

var mutex sync.Mutex //互斥锁(防止多个线程共同修改一个数据)

func add() {
	for i := 0; i < 5000; i++ {
		mutex.Lock() // 在数据改动的位置前加锁,其他线程碰到这个 会等待
		x = x + 1
		mutex.Unlock() // 当此线程对数据操作完成之后,解锁
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go add()
	}

	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 // 1000
	fmt.Printf("x : %d  耗时:%v \n", x, cost)
	fmt.Println("main quit")

}
