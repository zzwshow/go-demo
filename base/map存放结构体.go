package main

import "fmt"

//
type SS struct {
	Name string
}


func main(){
	ss := SS{Name:"zzw"}
	
	// 创建一个存放结构体类型的map
	
	ms := make(map[string]SS)
	ms["t1"] = ss
	
	fmt.Println(ms["t1"].Name)
	
}
