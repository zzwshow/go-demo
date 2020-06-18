package main

import (
	"fmt"
	"net"
	"os"
)

// ip







func main()  {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	
	name := os.Args[1]
	addr := net.ParseIP(name)  // 分析 命令行参数是否是合法IP，  返回addr 为：net.IP 类型
	fmt.Printf("%T\n",addr)
	
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
	
	

}
