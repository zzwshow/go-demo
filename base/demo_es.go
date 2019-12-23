package main

import "fmt"

func main() {
	var qq []string
	b := checkslice(qq)
	fmt.Println(b)
}

func checkslice(con []string) bool {
	fmt.Println("length:", len(con))
	if len(con) == 0 {
		return true
	}
	return false
}
