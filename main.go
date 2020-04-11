package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 判断目录或文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
		return false, err
	}
	return true, nil
}

// 复制已有的文件到新文件 不带缓冲区的
//func CopyFileToNewFile(sourceFile, destFile string) error {
//	data, err := ioutil.ReadFile(sourceFile)
//	if err != nil {
//		return err
//	}
//	err = ioutil.WriteFile(destFile, data, 0666)
//	if err != nil {
//		return err
//	}
//	return nil
//}

// 复制已有的文件到新文件 带缓冲区的
func CopyFileToNewFile_2(sourceFileName, destFileName string) error {
	srcFile, err := os.Open(sourceFileName)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	destFile, err := os.OpenFile(destFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(destFile)
	defer destFile.Close()
	_, err = io.Copy(writer, reader)
	writer.Flush()
	return err
}

// 读取文件内容
func GetFileContent(fileName string) (string, error) {
	// 1、带缓冲读文件内容
	// 2、直接读取文件内容
	// 先判断文件是否存在
	var s string
	bo, err := PathExists(fileName)
	if err != nil {
		return s, err
	}
	if !bo {
		return s, err
	}
	// 第一种带缓冲区读
	file, err := os.Open(fileName)
	if err != nil {
		return s, err
	}
	defer file.Close()
	var content []byte
	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		content = append(content, buf[:n]...)
		if err == io.EOF {
			break
		}
	}

	// 不带缓冲区
	// 一次性读取所有文件内容,适用于文件内容不大的情况
	// ioutil.ReadFile 返回两个值,一个是读取到的内容字节切片,一个是返回的错误信息
	// ioutil.ReadFile 不需要显示的close 文件, open 和 close 方法被封装在ReadFile 内部
	//content,err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	return s,err
	//}

	return string(content), err
}

// string写文件
func writeStringToFile(filepath, content string) error {
	//打开文件，没有则创建，有则append内容
	w1, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	_, err = w1.Write([]byte(content))
	if err != nil {
		return err
	}
	err = w1.Close()
	return err
}

func main() {
	sourcePath := "/Users/zzw/Desktop/eds.txt"
	//destFile := "/Users/zzw/Desktop/demo.sh"
	//fmt.Println("原文件：", sourcePath)
	//err := CopyFileToNewFile_2(sourcePath, destFile)
	//if err != nil {
	//	fmt.Println("复制文件异常....", err)
	//	return
	//}
	//fmt.Println("复制的新文件已经成功创建：", destFile)

	//bo, err := PathExists(sourcePath)
	//if err != nil {
	//	fmt.Println("判断路径是否存在时发生异常： ERROR：", err)
	//	return
	//}
	//if bo {
	//	fmt.Println("路径是存在的...")
	//	return
	//}
	//fmt.Println("路径不存在...")

	// 读取文件内容
	//content, err := GetFileContent(sourcePath)
	//if err != nil {
	//	fmt.Println("读取文件异常...", err)
	//	return
	//}
	//fmt.Println("文件内容是：\n", content)

	testStr := "asdfajdfjsljl"
	// 写string 到文件
	err := writeStringToFile(sourcePath, testStr)
	if err != nil {
		fmt.Println("写入文件失败...")
		return
	}
	fmt.Println("写入文件成功")

}
