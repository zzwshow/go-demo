package main

import (
	"fmt"
	"time"
)

// --- tasks ---
type Task struct {
	f func() error
}

// 创建一个任务
func NewTask(fc func() error) *Task {
	return &Task{f: fc}
}

// 任务的执行方法
func (t *Task) Execute() {
	t.f() // 调用任务函数
}

// ---- goroutine Pool ----
type Pool struct {
	// 对外的Task 入口
	EntryChannel chan *Task
	// 对内的Task队列
	JobsChannel chan *Task
	// 协程池中允许运行的最大worker数量
	workerNum int
}

// 创建协程池Pool 的函数
func NewPool(cap int) *Pool {
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		workerNum:    cap,
	}
	return &p
}

// 协程池创建一个worker，并且让这个worker执行任务
func (p *Pool) worker(workerId int) {
	// 一个具体的工作者(侦听JobsChannel管道中的任务)
	// 1、永久的从JobsChannel 中取任务并执行
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("Worker Id:", workerId, "任务已经执行完毕...")
	}
}

// 让协程池开始工作
func (p *Pool) run() {
	// 1 根据workerNum来创建需要几个worker来工作
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
	// 2 从EntryChannel中取任务，将取到的任务(这里可以做一些业务逻辑处理)，发送给JobChannel
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}

func taskA() error {
	fmt.Println("当前系统时间：", time.Now())
	return nil
}

// 主函数中测试，协程池的工作
func main() {
	// 1 创建一些任务
	t := NewTask(taskA)

	// 2 创建一个协程池Pool
	p := NewPool(4)

	// 3 将这些任务，传递给协程池Pool (必须要将入口放在go携程内，不然p.run永远也执行不到)
	go func() {
		for { //  死循环 一直往入口通道增加任务
			time.Sleep(1 * time.Second)
			p.EntryChannel <- t
		}
	}()

	// 4 启动协程池
	p.run()
}

//
