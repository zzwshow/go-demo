package main

import "fmt"

func main() {

	// ip := "172.16.0.1"
	// appname := "base"
	// checkurlP := "http://%s/galaxy-%s-business/static/app/common/healthCheck"
	// checkurl := fmt.Sprintf(checkurlP, ip, appname)
	// fmt.Println(checkurl)

	applist := make([]int64, 5)
	if applist[0] > 1 {
		fmt.Println("ok")
		return
	}
	fmt.Println("no")

}
