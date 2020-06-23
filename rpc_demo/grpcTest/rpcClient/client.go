package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

// 调用服务
func main() {
	// 1连接远程rpc服务
	rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("Client:", err)
	}
	// 2 调用远程方法
	// 定义接收服务端传回来的计算结果的变量
	ret := 0
	// 求面积
	err = rp.Call("Rect.Area", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("面积：", ret)

	// 求周长
	err = rp.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长：", ret)

}
