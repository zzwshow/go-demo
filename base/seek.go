package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	/*
		Seek(offset int64, whence int) (int64,error), 设置指针光标的位置
		第一个参数: 偏移量
		第二个参数: 如何设置
	*/
	fileName := "/home/zzw/Desktop/git/go-demo/aa.txt"
	file, err := os.OpenFile(fileName,os.O_RDWR,os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bs := []byte{0} //读取一个字节
	file.Read(bs)
	fmt.Println(string(bs))  //返回首字母a

	// 设置光标的位置
	file.Seek(4,io.SeekStart) //设置游标从第四个字符开始读取
	file.Read(bs)
	fmt.Println(string(bs))  //返回e



}

