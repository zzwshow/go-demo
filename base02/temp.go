package main

import (
	"fmt"
	"reflect"
)

func Sum(a int, more []int) int {
	for _, v := range more {
		a += v
	}
	return a
}

var funcs map[string]reflect.Value

func T(fName string, i interface{}) {
	fVal := reflect.ValueOf(i)
	funcs[fName] = fVal
	fmt.Println(fVal)
}

func main() {
	T("sum", Sum)
}
