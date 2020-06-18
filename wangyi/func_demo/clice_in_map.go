package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var s []map[string]int // 定义的是切片,但是切片内存放的是map
	s = make([]map[string]int,5,16)
	for index,val := range s{
		fmt.Printf("slice:[%d]=%v\n",index,val)
	}
	s[0] = make(map[string]int,16) // 初始化切片内的map
	s[0]["stu01"] = 1000
	s[0]["stu02"] = 2000
	s[0]["stu03"] = 3000
	s[0]["stu04"] = 4000

	for index,val := range s{
		fmt.Printf("slice:[%d]=%v\n",index,val)
	}
}
