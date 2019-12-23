package main

import (
	"fmt"
	"net"
)

// 服务端流程
/*
1. 使用net包,启动监听一个tcp server
2. 写一个死循环,一直监听客户端的链接
3. 编写读取客户端链接数据的函数,
4. 若第二步,客户端链接成功,就起一个单独的gorutine 处理这个链接,读取客户端数据之后,defer关闭当前链接
*/

func main() {
	fmt.Println("启动 tcp Server....")
	listen, err := net.Listen("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Printf("listen failed ,err:%v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err %v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read data failed...")
			break
		}
		str := string(buf[:n])
		fmt.Printf("recv from client, data:%v\n", str)

	}
}
