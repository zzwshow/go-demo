package main

import (
	"fmt"
	"strings"
)

//统计没个单词出现的次数

func statWordCount(s string) map[string]int {

	var result map[string]int = make(map[string]int)

	words := strings.Split(s," ")
	fmt.Println(words)
	for _,v := range words{  // 循环这个切片
		value, ok := result[v]  //判断切片内的值是否已经存在map 中了,
		if !ok{                // 如果不存在,就新建,并讲value 至为1
			result[v] = 1
		} else {               // 如果已经存在,就讲原来的value 值加1
			result[v] = value +1
		}
	}
	return result

}


func main() {
	str := "Where did you eat your meal today and did you drink?"

	count_result := statWordCount(str)
	fmt.Println(count_result)


}
