package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2
	// 接口类型的变量----> 反射类型对象
	value := reflect.ValueOf(num)

	//反射类型对象-----> 接口类型变量
	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)


 // 指针
 pointer := reflect.ValueOf(&num)
 convertPointer := pointer.Interface().(*float64)
 fmt.Println(convertPointer)


}
