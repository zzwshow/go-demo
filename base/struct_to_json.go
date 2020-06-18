package main

#import (
#	"encoding/json"
#	"fmt"
#)
#
#type IT struct {
#	Company string
#	Subjects []string
#	IsOk bool
#	Price float64
#}
#
#
#func main(){
#	//定义结构体变量并初始化
#	s := IT{"itcast",[]string{"GO","PYTHON","java"},false,66.66}
#	// 编码先生成切片(不含缩进)
#	//buf,err := json.Marshal(s)
#	buf,err := json.MarshalIndent(s,""," ")//带缩进格式化
#
#
#	if err !=nil{
#		fmt.Println("err = ",err)
#	}
#	fmt.Printf("%T\n",buf)
#	fmt.Println(string(buf))
#
#}

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Logintime time.Time
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func main() {
	user := User{5, "zhangsan", "pwd", time.Now()}
	data := Struct2Map(user)
	fmt.Println(data)
}
