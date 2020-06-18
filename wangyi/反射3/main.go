package main

import (
	"fmt"
	"reflect"
)

// 未知类型的反射

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Sex string `json:"sex"`
}

func (p Person) Say(msg string){
	fmt.Println("hello, ",msg)
}

func (p Person) PrintInfo(){
	fmt.Printf("姓名: %s 年龄:%d 性别:%s",p.Name,p.Age,p.Sex)
}

func main() {
	p1 := Person{"王二够",31,"男"}
	GetMessage(p1)
}



// 获取input 信息
func GetMessage(input interface{}){
	getType := reflect.TypeOf(input) 	// 通过反射获取传进来的类型
	fmt.Println("get Type is :",getType)  //main.Person
	fmt.Println("get kind is :",getType.Kind()) //种类是 结构体 struct

	// 获取值
	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is : ",getValue) // {王二够 31 男}

	//获取字段
	/*
		step1: 先获取Type 对象,reflect.Type
				解构体类型,有NumField() 返回整数,有多少个字段
				Field(index)
		step2: 通过Filed()获取每一个Field字段
		step3: Interface() 得到对应的value
	*/
	for i:=0; i< getType.NumField();i++{
		filed := getType.Field(i)
		value := getValue.Field(i).Interface() //根据下标获取字段的数值
		fmt.Printf("字段的名称: %s 字段的类型:%s 字段数值:%v 字段的Tag:%v\n",filed.Name,filed.Type,value,filed.Tag.Get("json"))
	}

	// 获取方法
	for i:=0;i<getType.NumMethod();i++{
		method := getType.Method(i)
		fmt.Printf("方法名: %s  方法类型: %v\n",method.Name,method.Type)
	}


}


