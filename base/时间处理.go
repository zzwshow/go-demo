package main

import (
	"fmt"
	"time"
)

func main() {

	// var TestTime time.Time
	// TestTime = time.Time{}
	const DefaultTimeFormat = "2006-01-02 15:04:05"
	TestTime := time.Now()
	old7Time := TestTime.AddDate(0, 0, -7)

	fmt.Println(old7Time.Format(DefaultTimeFormat))

	// const DefaultTimeFormat_2 = "2006-01"
	// fmt.Println(TestTime.Format(DefaultTimeFormat_2))

	// if TestTime.IsZero() {
	// 	fmt.Println("此时间是无效的")
	// } else {
	// 	fmt.Println("这是个有效时间")
	// }

	fmt.Println()

}
