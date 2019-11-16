package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// tcp rpc client

type Args struct {
	A,B int
}

type Quotient struct {
	Quo, Rem int
}



func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("Usage : ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]                           // 通过命令行接收要连接的rpc server 主机地址
	client, err := rpc.Dial("tcp", serverAddress+":1234") // 注意：和http的客户端代码对比，唯一的区别一个是DialHTTP，一个是Dial(tcp)，其他处理一模一样。
	if err != nil {
		log.Fatal(err.Error())
	}
	
	args := Args{17, 8}
	var reply int
	err = client.Call("WorkRPC.Multiply", args, &reply) // 调用远程方法处理数据（计算7*8 的乘积）
	if err != nil {
		fmt.Println("rpc 调用 Multiply 方法失败：", err.Error())
	}
	fmt.Printf("WorkRPC: %d*%d=%d\n", args.A, args.B, reply)
	
	var quo Quotient
	err = client.Call("WorkRPC.Divide", args, &quo)
	if err != nil {
		fmt.Println("rpc 调用 Divide 方法失败：", err.Error())
	}
	fmt.Printf("WorkRPC: %d/%d=%d remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)
	
	/*
		结果：
			WorkRPC: 17*8=136
			WorkRPC: 17/8=2 remainder 1
	
	*/
	
}