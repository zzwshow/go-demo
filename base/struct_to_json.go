package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string
	Subjects []string
	IsOk bool
	Price float64
}


func main(){
	//定义结构体变量并初始化
	s := IT{"itcast",[]string{"GO","PYTHON","java"},false,66.66}
	// 编码先生成切片(不含缩进)
	//buf,err := json.Marshal(s)
	buf,err := json.MarshalIndent(s,""," ")//带缩进格式化


	if err !=nil{
		fmt.Println("err = ",err)
	}
	fmt.Printf("%T\n",buf)
	fmt.Println(string(buf))

}

