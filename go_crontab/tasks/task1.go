package tasks

import (
	"fmt"

)

type HealthCheck struct {
}

func (h *HealthCheck) Run(parameter string) (str string,err error){
	fmt.Println("我在健康检查任务....", parameter)
	return parameter,nil
}




