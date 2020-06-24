package main

import "fmt"

func main() {
	xx := "[CPU]Processor Util (1 min average)>90%"
	bb := "[CPU]Processor Util (1 min average)>90%"
	if xx == bb {
		fmt.Println("ok")
		return
	}
	fmt.Println("no")

}
