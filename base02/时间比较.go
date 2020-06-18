package main

import (
	"fmt"
	"time"
)

func main() {
	format := "2006-01-02 15:04:05"
	now := time.Now()
	a, _ := time.Parse(format, "2020-05-06 09:31:30")
	//b, _ := time.Parse(format, "2015-03-10 16:00:00")
	//fmt.Println("Now:", now.Format(format), "\na:", a.Format(format), "\nb:", b.Format(format))
	//fmt.Println("-----  a > now > b")
	t7, _ := time.ParseInLocation(time.RFC3339, "2020-05-13T08:24:10.877457147Z", time.Local)
	fmt.Println(t7.Format(format))
	fmt.Println("t7:", t7) // 输出：t7: 2018-10-01 16:27:00 +0800 CST
}
