package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	/*
		bufio: 高小io 读写
		buffer 缓存
	讲io包下的Reader,write对象进行包装,带缓存的包装,提高读写效率
	ReadBytes()
	ReadString() //读取字符串(传入分隔符,比如\n,读取到换行后结束)
	ReadLine()
	*/

	fileName := "/home/zzw/Desktop/git/go-demo/aa.txt"
	file,err := os.Open(fileName) // 只要获取文件的打开对象即可,读写操作使用bufio包
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建Reader 对象
	b1 := bufio.NewReader(file)
	p := make([]byte,1024) // 设置每次读取多少字节
	n1,err := b1.Read(p) //返回读取到的字符数,和错误
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))


}
