package http_rpc_server

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

// http rpc 服务端
/*
正确的RPC函数格式如下：
func (t *T) MethodName(argType T1, replyType *T2) error
*/

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
	w := new(WorkRPC)   // 创建结构体指针，分配内存
	e := rpc.Register(w) // 注册
	if e != nil{
		fmt.Println("rpc server register error,",e.Error())
	}
	rpc.HandleHTTP()    // 已http 协议注册服务
	
	err := http.ListenAndServe(":1234",nil)
	if err !=nil {
		fmt.Println(err.Error())
	}

}
