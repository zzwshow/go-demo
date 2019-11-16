package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

// tcp rpc server



//定义客户端需要传过来处理的参数结构体
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type WorkRPC struct {  // rpc 接口结构体
}

func (w *WorkRPC) Multiply(args *Args, reply *int) error {  // 为接口结构体定义计算乘积的方法
	*reply = args.A * args.B
	return nil
}

func (w WorkRPC) Divide(args *Args, quo *Quotient) error {  // 为接口结构体定义取商和余数的方法
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B  // 计算两个数的商
	quo.Rem = args.A % args.B  // 计算两个数取余
	return nil
}

func main()  {
	w := new(WorkRPC)
	err := rpc.Register(w)  // 注册rpc
	checkError(err)
	
	tcpAddr,err := net.ResolveTCPAddr("tcp",":1234")  // 解析tcp 绑定的地址
	checkError(err)
	
	listener,err := net.ListenTCP("tcp",tcpAddr)  // 监听此tcp 地址
	checkError(err)
	
	for {
		conn,err := listener.Accept()  //接受在侦听器接口中实现接受方法
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)   // cpc 接收连接
		// 这个代码和http的服务器相比，不同在于:在此处我们采用了TCP协议，
		// 然后需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc来处理。
	}

}

// 检查err
func checkError(err error){
	if err !=nil{
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
