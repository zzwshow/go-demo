package main

import (
	"fmt"
	"reflect"
)
type test struct {
	aa int
}

type Person struct {
	Nmae string
	Age int
	Test test
}


func main(){
	ss := Person{Nmae:"zzw",Age:29}
	
	v:= reflect.ValueOf(ss)  // 获取结构体内的字段值
	count := v.NumField()    // 统计结构体内的字段数量
	fmt.Println(v)
	fmt.Println(count)
}
