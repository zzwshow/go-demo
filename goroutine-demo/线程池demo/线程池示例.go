package 线程池demo

import (
	"fmt"
	"time"
)

type Pool struct { // 定义线程池
	Queue         chan func() error // 队列 存放执行函数的管道，并返回错误
	RuntineNumber int               // 控制 线程队列缓存数量
	Total         int               // 记录 全部任务数

	Result         chan error // 结果 存放执行结果的管道
	FinishCallback func()     // 回调 函数类型 没有返回值
}

// 初始化
func (self *Pool) Init(runtineNumber int, total int) { // 初始化结构体
	self.Queue = make(chan func() error, total)
	self.RuntineNumber = runtineNumber
	self.Total = total
	self.Result = make(chan error, total)
}

func (self *Pool) Start() {
	// 开起 number 个goroutine
	for i := 0; i < self.RuntineNumber; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}
				err := task() // 任务执行后
				self.Result <- err
			}
		}()
	}

	// 获取每个任务的处理结果
	for j := 0; j < self.RuntineNumber; j++ {
		res, ok := <-self.Result
		if !ok {
			break
		}
		if res != nil {
			fmt.Println(res)
		}
	}

	// 结束回调函数
	if self.FinishCallback != nil {
		self.FinishCallback()
	}
}

func (self *Pool) AddTask(task func() error) {
	self.Queue <- task
}

func (self *Pool) SetFinishCallback(fun func()) {
	self.FinishCallback = fun
}

// 关闭
func (self *Pool) Stop() {
	close(self.Queue)
	close(self.Result)
}

func main() {
	var p Pool
	url := []string{"11111", "22222", "33333", "444444", "55555", "66666", "77777", "88888", "999999"}
	p.Init(9, len(url))

	for i := range url {
		u := url[i]
		p.AddTask(func() error {
			return Download(u)
		})
	}

	p.SetFinishCallback(DownloadFinish)
	p.Start()
	p.Stop()
}

func Download(url string) error {
	time.Sleep(1 * time.Second)
	fmt.Println("Download" + url)
	return nil
}

func DownloadFinish() {
	fmt.Println("Download finish。。。")
}
