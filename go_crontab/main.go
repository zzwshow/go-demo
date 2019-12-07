package main
// cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"time"
	"go_crontab/tasks"
)


func createJob(taskname string, parameter string) cron.FuncJob {
	handler := tasks.CreateHandler(taskname)
	taskFunc := func() {
		handler.Run(parameter)
	}
	return taskFunc
}



func main(){
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	cron.WithLocation(nyc)
	c := cron.New(cron.WithSeconds())
	c.Start()
	defer c.Stop()
	
	execTime := "*/4 * * * * *"
	parameter := "我是任务参数"
	taskname := "healthcheck"
	t_func :=createJob(taskname,parameter)
	
	
	t_id,err := c.AddJob(execTime,t_func)
	checkError(err)
	fmt.Println("param:",t_id)
	select {}
}


func checkError(err error) {
	if err != nil{
		fmt.Println("发生错误：",err.Error())
		os.Exit(1)
	}
}



