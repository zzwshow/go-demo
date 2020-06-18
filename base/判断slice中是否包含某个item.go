package main

import (
	"fmt"
	"reflect"
)

// 判断int clice 中是否包含某个id  (只能对一种数据类型判断)
func IsExistInArray(value int, array []int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// 判断slice中是否存在某个item  (可对多种数据类型进行判断)
func IsExistItem2(value interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func main() {
	sss := []string{"a", "b", "c", "d"}
	item := "a"
	b := IsExistItem2(item, sss)
	fmt.Println(b)

}
