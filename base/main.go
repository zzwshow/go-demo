package main

import (
	"fmt"
	"reflect"
)

func main(){

	page := 0
	limit := 0
	if reflect.ValueOf(page).IsZero() {
		fmt.Println("true")
	}
	if reflect.ValueOf(limit).IsZero() {
		fmt.Println("false")
	}

}
