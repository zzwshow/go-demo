package main

import (
	"fmt"
	"strings"
)

//  此函数不能截取中文 （先获取子串前的索引位）
func UnicodeIndex(str,substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str,substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}
	return result
}

func main()  {
	Str := "galaxy-base-business"
	//index := UnicodeIndex(Str,".")
	fmt.Println(Str[0:0])
}
