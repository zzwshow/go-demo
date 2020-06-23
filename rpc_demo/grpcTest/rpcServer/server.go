package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// 服务端，求矩形面积和周长

// 声明一个矩形对象
type Rect struct {
}

// 声明参数结构体
type Params struct {
	Width, Height int
}

// 定义求面积的方法
func (r *Rect) Area(p Params, result *int) error {
	*result = p.Height * p.Width
	return nil
}

// 定义求周长的方法
func (r *Rect) Perimeter(p Params, result *int) error {
	*result = (p.Width + p.Height) * 2
	return nil
}

func main() {
	// 1 注册服务
	rect := new(Rect)
	rpc.Register(rect)
	// 2 把处理服务绑定到http协议上
	rpc.HandleHTTP()
	// 3 监听服务，等待客户端调用求面积和周长的方法
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Server :", err)
	}
}
