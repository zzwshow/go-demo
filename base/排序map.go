package main

import (
	"fmt"
	"sort"
)

func main() {
	scene := make(map[string]int)
	//准备数据
	scene["route"] = 66
	scene["brec"] = 100
	scene["kkasjdf"] = 10
	
	// 声明一个切片保存map 数据
	var sceneList []string
	
	//将map数据遍历复制到切片中
	for k := range scene{
		sceneList = append(sceneList,k)
	}
	
	//对切片进行排序
	sort.Strings(sceneList)
	
	fmt.Println(sceneList) //[brec kkasjdf route]
	
	
	
	
}
