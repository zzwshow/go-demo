package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	Home string
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func main() {
	student := Student{"小明", 20, ""}
	mapData := StructToMap(student)
	fmt.Printf("%T\n, student: %#v", mapData, mapData)
}
