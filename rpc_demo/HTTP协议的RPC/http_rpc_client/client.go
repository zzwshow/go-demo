package http_rpc_client

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// rpc HTTP协议 调用方，客户端


type Args struct {
	A,B int
}

type Quotient struct {
	Quo, Rem int
}



func main(){
	
	if len(os.Args) < 2 {
		fmt.Println("Usage : ",os.Args[0],"server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]         // 通过命令行接收要连接的rpc server 主机地址
	client, err := rpc.DialHTTP("tcp",serverAddress+":1234")  // 建立HTTP连接
	if err != nil{
		log.Fatal(err.Error())
	}
	
	args := Args{17,8}
	var reply int
	err = client.Call("WorkRPC.Multiply",args,&reply)  // 调用远程方法处理数据（计算7*8 的乘积）
	if err!=nil{
		fmt.Println("rpc 调用 Multiply 方法失败：",err.Error())
	}
	fmt.Printf("WorkRPC: %d*%d=%d\n", args.A, args.B, reply)
	
	
	var quo Quotient
	err =client.Call("WorkRPC.Divide",args,&quo)
	if err != nil{
		fmt.Println("rpc 调用 Divide 方法失败：",err.Error())
	}
	fmt.Printf("WorkRPC: %d/%d=%d remainder %d\n", args.A, args.B, quo.Quo, quo.Rem)
	
	
	/*
	结果如下：
	WorkRPC: 17*8=136
	WorkRPC: 17/8=2 remainder 1
	
	
	总结
	通过上面的调用可以看到参数和返回值是我们定义的struct类型，在服务端我们把它们当做调用函数的参数的类型，
	在客户端作为client.Call的第2，3两个参数的类型。客户端最重要的就是这个Call函数，它有3个参数，第1个要调用的函数的名字，
	第2个是要传递的参数，第3个要返回的参数(注意是指针类型)，通过上面的代码例子我们可以发现，使用Go的RPC实现相当的简单，方便。
	*/
	
}
