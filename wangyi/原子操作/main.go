package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int32
var wg sync.WaitGroup

//var mutex sync.Mutex //互斥锁(防止多个线程共同修改一个数据)

func add() {
	for i := 0; i < 5000; i++ {
		//mutex.Lock() // 在数据改动的位置前加锁,其他线程碰到这个 会等待
		//x = x + 1
		//mutex.Unlock() // 当此线程对数据操作完成之后,解锁

		atomic.AddInt32(&x, 1) // 取地址符,加1操作,比互斥锁性能要高
	}
	wg.Done()
}

func main() {
	wg.Add(3)
	go add()
	go add()
	go add()

	wg.Wait()
	fmt.Printf("x : %d  \n", x)
	fmt.Println("main quit")

}
