package main

import "fmt"

func main() {
	// var  a []int
	// a = append(a,1)
	// a = append(a,2,3,4,5)
	// a = append(a,[]int{6,7,8}...)  // 追加一个切片, 切片需要解包 //[1 2 3 4 5 6 7 8]
	// fmt.Println(a)
	//
	//
	// var b []int
	// b = append(b,1)
	// b = append([]int{0}, b...)  //[0 1] 在切片的起始位置添加一个元素，切片需要解包
	// b = append([]int{-3,-2,-1}, b...) // 在切片的起始位置添加一个切片 [-3 -2 -1 0 1]
	// fmt.Println(b)
	
	
	//删除
	aa := []int{1,2,3}
	fmt.Println(aa)
	aa = aa[1:]
	fmt.Println(aa)
	
}
