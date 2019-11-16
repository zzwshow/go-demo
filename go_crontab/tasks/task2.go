package tasks

import "fmt"

type AppYnc struct {
	App string
}

func (a AppYnc) Run(){
	fmt.Println("我是app信息同步任务....")
}
