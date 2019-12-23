package main

import "fmt"
//4,3,2,5,6,7,8,1
func insert_sort(a [8]int) [8]int {
	for i := 1; i<len(a); i++{ //i=1
		for j := i; j>0; j--{  // j=1
			if a[j] < a[j-1]{
				a[j],a[j-1] = a[j-1],a[j]
			} else {
				break
			}
		}
	}
	return a
}


func main(){
	var i [8]int = [8]int{4,3,2,5,6,7,8,1}
	sort_i := insert_sort(i)
	fmt.Println(i)   //因为i传递给insert_sort 函数中是值传递,是传递i的copy,所以排序后不会影响i原本的值
	fmt.Println(sort_i)

}


