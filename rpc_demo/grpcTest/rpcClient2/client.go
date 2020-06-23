package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// 声明接收的参数结构体
type ArithRequest struct {
	A, B int
}

// 声明返回给客户端的数据结构体
type ArithResponse struct {
	// 乘机
	Pro int
	// 商
	Quo int
	// 余数
	Rem int
}

func main() {
	// 连接远程RPC
	// conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8082")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8082")

	if err != nil {
		log.Fatal(err)
	}
	reqParam := ArithRequest{A: 10, B: 3}
	response := new(ArithResponse)
	err = conn.Call("Arith.Multiply", reqParam, response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("乘机：", response.Pro)
	fmt.Println("商：", response.Rem)
	fmt.Println("余数：", response.Quo)

}
