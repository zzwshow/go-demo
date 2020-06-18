package main

import "fmt"

func main() {
	var a = make([]string,5,10) //长度为5容量为10 的切片
	for i:=0;i<10;i++{
		a = append(a,fmt.Sprintf("%v",i))
		// 将int型格式化为字符串

	}
	fmt.Println(a)

	/*
	[     0 1 2 3 4 5 6 7 8 9], 前面为5个空字符串,append 从第五个位置开始添加

	*/


}
