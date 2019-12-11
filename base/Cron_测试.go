package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type Hello struct {
	Str string
}

func(h Hello) Run(){
	fmt.Println("正在执行任务...",h.Str)
}

func main(){
	fmt.Println("Starting。。。")
	
	c := cron.New()   // 实例化定时任务对象
	h := Hello{Str:"zzw"}   // 实例化结构体
	spec := "*/2 * * * * *"
	
	c.AddJob("")
	
	
	
}
