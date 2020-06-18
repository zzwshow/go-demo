package main

import (
	"fmt"
	"sync"
)

/*
由于goroutine是异步执行的，那很有可能出现主程序退出时还有goroutine没有执行完，此时goroutine也会跟着退出。
此时如果想等到所有goroutine任务执行完毕才退出，go提供了sync包和channel来解决同步问题，当然如果你能预测每个goroutine执行的时间，
你还可以通过time.Sleep方式等待所有的groutine执行完成以后在退出程序(如上面的列子)。
*/
//使用sync包同步goroutine
/*
 sync大致实现方式:
 	WaitGroup 等待一组goroutinue执行完毕. 主程序调用 Add 添加等待的goroutinue数量. 每个goroutinue在执行结束时调用 Done ，
 此时等待队列数量减1.，主程序通过Wait阻塞，直到等待队列为0.
*/


func cal(a,b int, n *sync.WaitGroup) {  //这里接收指针参数，因为我们在函数内需要改变WaitGroup的计数的值
	defer n.Done() // 此 goroutinue完成后, WaitGroup的计数-1 (或者不用defer，放在函数最后一行)
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
}

func main(){
	var my_waitGroup sync.WaitGroup  // 声明一个WaitGroup 变量
	for i :=0; i<10; i++ {
		my_waitGroup.Add(1)          // 增加计数器值（需与启动动goroutine 数量一致）
		go cal(i,i+1, &my_waitGroup)
	}
	
	my_waitGroup.Wait() // 让主程序等待所有goroutine 执行完毕后在退出

}
