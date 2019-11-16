package main
// cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"time"
	
	"go_crontab/tasks"
)

func main(){
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	cron.WithLocation(nyc)
	c := cron.New(cron.WithSeconds())
	c.Start()
	defer c.Stop()
	
	execTime := "*/4 * * * * *"
	tasks.InitTaskList()
	
	// t_id,err := c.AddJob(execTime,tmp)
	// checkError(err)
	// fmt.Println("param:",t_id)
	select {}
}


func checkError(err error) {
	if err != nil{
		fmt.Println("发生错误：",err.Error())
		os.Exit(1)
	}
}
