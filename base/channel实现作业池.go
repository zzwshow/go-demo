package main

import "fmt"

//channel实现作业池

// 创建三个channel，一个channel用于接受任务，一个channel用于保持结果，还有个channel用于决定程序退出的时候。


func Task(taskChan,resultChan chan int,exitChan chan bool) {
	defer func() {
		err := recover()
		if err!= nil{
			fmt.Println("Task func 执行错误...",err)
			return
		}
	}()
	
	for t := range taskChan {
		fmt.Printf("Do Task is :%d\n",t)
		resultChan <- (t + 1)  // 将读取结果并+1 后放入结果管道
	}
	exitChan <- true // 上面函数执行完 给主协程发送退出信号
	
}


func main(){
	
	taskChan := make(chan int, 20)
	resultChan := make(chan int, 20)
	exitChan := make(chan bool, 5)
	
	go func() {           				//启动一个匿名函数的goroutine 来产生任务，执行完后关闭管道 (产生任务，完成后关闭管道)
		for i :=0; i<10; i++{
			taskChan <- i
		}
		close(taskChan)
	}()
	
	for j:=0; j<5; j++{   				// 启动5个goroutine 去执行任务  （启动5个goroutine 去执行任务）
		go Task(taskChan,resultChan,exitChan)
	}
	
	go func() {  // 等待5个goroutine 结束  						（接收 退出信号）
		for i:=0; i<5; i++{
			<-exitChan    // 这里使用管道阻塞 来等待goroutine 都执行完毕在退出
		}
		close(resultChan)  // 当5个任务跑完以后关闭管道
		close(exitChan)
	}()
	
	
	//  从结果管到内获取 数据
	for v := range resultChan{
		fmt.Println("任务处理结果是：",v)
	}
	
}






