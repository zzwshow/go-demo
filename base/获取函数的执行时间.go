package main

import (
	"fmt"
	"time"
)

//最简单的办法就是在函数执行之前设置一个起始时间，并在函数运行结束时获取从起始时间到现在所经过的时间间隔，就可以得出函数运行所消耗的具体时间。

func test_1() {
	start := time.Now() //获取当前时间
	sum := 0
	for i :=0; i<50000000; i++{
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：",elapsed)
}


func main() {
	
	test_1()


	
}
