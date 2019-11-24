package main

import "fmt"

func Adder() func(int) int{
	var x int
	return func(d int) int {
		x += d
		return x
	}
}


func main(){
	a := Adder()
	fmt.Println(a(1))
	fmt.Println(a(100))
	fmt.Println(a(200))
}

/*
1
101
301

*/