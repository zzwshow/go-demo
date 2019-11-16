package tasks

import "fmt"

type HealthCheck struct {
	Parameter string
}

func (h HealthCheck) Run(){
	fmt.Println("我在健康检查任务....", h.Parameter)
}


