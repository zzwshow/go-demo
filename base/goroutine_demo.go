package main

import (
	"fmt"
	"runtime"
	"time"
)


func Cal(a int,b int) {
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
}

func main()  {
		// 获取主机逻辑CPU个数
		cpuNum := runtime.NumCPU()
		fmt.Println("主机逻辑CPU个数：",cpuNum)
		// runtime.GOMAXPROCS(2) // 设置最大使用的逻辑CPU 数量，1.8 以后默认使用全部！
		
		for i := 0; i<10; i++{
			go Cal(1,i+1)  //启动10 个goroutine 去执行
		}
		
		time.Sleep(time.Second * 2) // 防止goroutine 没有执行完，我们把主协程睡眠两秒
		
	/* 	结果如下：
		1 + 1 = 2
		1 + 6 = 7
		1 + 7 = 8
		1 + 4 = 5
		1 + 2 = 3
		1 + 8 = 9
		1 + 5 = 6
		1 + 9 = 10
		1 + 10 = 11
		1 + 3 = 4
	可以看到，结果并不是按照顺序打印的，因为10 goroutine 的执行先后是runtime自己控制的
	*/
	
	
}


