package tasks

import (
	"fmt"
)

type AppYnc struct {
}

func (a *AppYnc) Run(parameter string) (str string, err error) {
	fmt.Println("我是app信息同步任务....",parameter)
	return parameter,nil
}



