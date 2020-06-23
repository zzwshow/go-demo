package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Arith struct {
}

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

// 乘法运算
func (a *Arith) Multiply(req ArithRequest, res *ArithResponse) (err error) {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Pro = req.A * req.B
	// 商
	res.Quo = req.A / req.B
	// 余数
	res.Rem = req.A % req.B
	return nil
}

//func main() {
//	// 注册服务
//	rpc.Register(new(Arith))
//	// 采用HTTP作为载体
//	rpc.HandleHTTP()
//	err := http.ListenAndServe(":8082", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}

//jsonRPC编码方式
func main() {
	// 注册服务
	rpc.Register(new(Arith))
	// 监听服务
	lis, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		// 每个连接都创建一个协程来跑
		go func(conn net.Conn) {
			fmt.Println("new a client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
