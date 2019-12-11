package main

import (
	"fmt"
	"time"
)

func sum(num int) (total time.Duration) {
	startTime := time.Now()
	
	for i:=0; i<num; i++{
		time.Sleep(time.Millisecond * 1)
	}
	endingTime := time.Now()
	
	totalTime := endingTime.Sub(startTime)
	return totalTime
}


func main() {
	
	var num int = 100
	totalTime := sum(num)
	fmt.Printf("%T\n",totalTime)
	fmt.Println("总耗时：",totalTime.String())
	
	// var float64_Minutes float64 = totalTime.Seconds()
	// fmt.Println("总耗时：",float64_Minutes)
	
}
