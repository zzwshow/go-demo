package rpc

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

// 测试连接读写
func TestSessionReadWrite(t *testing.T) {
	// 定义监听Ip 和端口
	addr := "127.0.0.1:8083"
	// 定义要传输的数据
	Mydata := "hello world"
	// 等待组
	wg := sync.WaitGroup{}
	// 增加2个计数器
	wg.Add(2)
	// 协程1个读 ，1个写
	// 写协程
	go func() {
		defer wg.Done()
		// 创建tcp 连接
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, _ := lis.Accept()
		s := NewSession(conn)
		// 通过会话写入数据
		err = s.Write([]byte(Mydata))
		if err != nil {
			t.Fatal(err)
		}
	}()
	// 读协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := NewSession(conn)
		data, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != Mydata {
			t.Fatal(err)
		}
		fmt.Println("data:", string(data))
	}()
	wg.Wait()
}
