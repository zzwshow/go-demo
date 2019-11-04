package main

import (
	"fmt"
	"math/rand"
)

// channel 练习
// 生产者和消费者模型
// 使用goroutine 和channel 实现

// 生产者: 产生随机数
// 消费者: 计算每个随机数的每个位的数字的和

// 声明一个全局通道
var itemChan chan *Item   // 通道内存放的是结构体指针
var resultChan chan *result

type Item struct { //生产者使用的结构体
	id int64
	num int64
}

type result struct { //消费者使用的结构体
	Item *Item
	sum int64
}

// 生产者
func producer(ch chan *Item){
	// 1. 生成随机数,2. 将随机数发送到通道内
	var id int64
	for {      // 循环往通道内写随机数(死循环)
		id++
		number := rand.Int63() // int64 正整数
		tmp := &Item{
			id:  id,
			num: number,
		}
		ch <- tmp  // 将指针类型的结构体(随机数)放入通道内
	}
}

//消费者(ch 通道内是生产者生产的随机数,resultChan通道是消费者接收以后存放的通道)
func consumer(ch chan *Item,resultChan chan *result){
	for tmp :=  range ch {  // 不断的从生产者的通道内取值,计算后放入到消费者的通道内
		// 从通道取值,传递给计算函数
		sum := calc(tmp.num)
		// 构造result 结构体
		retObj := &result{
			Item: tmp,
			sum:  sum,
		}
		resultChan <-retObj  // 将消费者的数据放入到消费者的通道内
	}
}
// 计算一个数字每个位的和
func calc(num int64) int64 {
	var sum int64
	for num > 0{
		sum = sum + num%10
		num = num/10
	}
	return sum
}


//打印结果(从消费者的通道里取出 结果)
func printResult(resultChan chan *result){
	for ret:= range resultChan{
		fmt.Printf("id:%v,num:%v,sum:%v\n",ret.Item.id,ret.Item.num,ret.sum)
	}
}
// 控制启动多少个消费者goroutine
func startWorker(n int,ch chan *Item,resultChan chan *result){
	for i:=0;i<n;i++{
		go consumer(ch,resultChan)
	}
}

func main(){
	//初始化通道
	itemChan = make(chan *Item,100)
	resultChan = make(chan *result,100)
	go producer(itemChan) // 讲通道传给生产者函数


	startWorker(10,itemChan,resultChan) //启动10个goroutine消费者

	// 打印结果
	printResult(resultChan)

	// 打印随机数
	//rand.Seed(time.Now().Unix()) // 只有添加了随机数种子,才会在每次运行都产生随机数
	//ret := rand.Int63() //iint64 正整数
	//fmt.Println(ret)
	//ret2 := rand.Intn(101) //取值范围 [1.101]
	//fmt.Println(ret2)

}

