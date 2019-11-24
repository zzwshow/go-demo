package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Sex int
	Age int
}





func main() {
	var s Student

	v := reflect.ValueOf(s)  // 通过反射拿到结构体内成员的值  偶偶他:v:{ 0 0}
	t := v.Type()            // out:t:main.Student
	fmt.Printf("v:%v  t:%v",v,t)

	k := t.Kind()
	switch (k) {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Struct:
		fmt.Println("struct")
	}


}
