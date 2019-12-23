package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp 客户端
/*
1. 建立与服务端的链接
2. 进行数据收发
3. 关闭链接
*/
func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:20000")
	if err != nil {
		fmt.Println("connect TCP server failed")
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		trimmedlnput := strings.Trim(input, "\r\n")
		if trimmedlnput == "Q" || trimmedlnput == "exit" {
			return
		}
		_, err := conn.Write([]byte(trimmedlnput))
		if err != nil {
			return
		}
	}

}
