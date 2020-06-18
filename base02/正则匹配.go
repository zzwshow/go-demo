package main

import (
	"fmt"
	"regexp"
)

func main() {
	s1 := "[数据库]10.205.64.31 mycat 80220 实例 jmx 不可用，当前值为 10 days, 23:10:47"
	reg := regexp.MustCompile(`\d{4,5}`)
	fmt.Println(reg.FindAllString(s1, -1)[0])
}
