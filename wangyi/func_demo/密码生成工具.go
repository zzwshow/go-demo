package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// flag
// 密码生成器

var (
	length int
	charset string
)

const (
	NumStr = "0123456789"
	CharSet = "asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP"
	SpecSet = "+=-@#$%^&*()!"
)


func parseArgs(){  //解析命令行参数的函数
	//定义第一个选项
	flag.IntVar(&length,"l",16,"-l 指定要生成密码的长度默认是16位")
	// 第二个选项
	flag.StringVar(&charset,"t","num",
		`-t 制定生成密码的字符集,
				num:只使用数字[0-9],
				char:只使用字母[a-zA-Z],
				mix:使用数字和字母,
				advance:使用数字,字母以及特殊字符`,
	)
	flag.Parse()  //解析参数
}


func generatePassword() string {
	var password []byte = make([]byte,length,length)  //定义一个存放密码的切片,长度为0,容量是用户输入的长度

	var sourcesStr string
	if charset == "num" {
		sourcesStr = NumStr
	} else if charset == "char" {
		sourcesStr = CharSet
	} else if charset == "mix" {
		sourcesStr = fmt.Sprintf("%s%s",NumStr,CharSet)
	} else if charset == "advance" {
		sourcesStr = fmt.Sprintf("%s%s%s",NumStr,CharSet,SpecSet)
	} else {
		sourcesStr = NumStr
	}
	//fmt.Printf("用户选择的是: %s, 包含字符:%v\n",charset,sourcesStr)


	for i:=0;i<length;i++{   //根据用户选择的字符,随机生成密码(根据sourcesStr,随机生成下标,来生成)
		index := rand.Intn(len(sourcesStr)-1)
		password[i] = sourcesStr[index]  //根据随机下标取出对应的字符,赋值给密码切片
	}
	return string(password)
}



func main() {
	rand.Seed(time.Now().UnixNano())  //定义随机数种子

	parseArgs() // 接受用户命令行输入
	//fmt.Printf("length:%d charset:%s\n",length,charset)

	//生成密码
	password := generatePassword()
	fmt.Println("password:",password)
}
