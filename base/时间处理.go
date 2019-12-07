package main

import (
	"fmt"
	"time"
)

func main() {
	
	var TestTime time.Time
	TestTime = time.Time{}
	// TestTime = time.Now()
	
	
	if TestTime.IsZero(){
		fmt.Println("此时间是无效的")
	}else {
		fmt.Println("这是个有效时间")
	}
	
}
