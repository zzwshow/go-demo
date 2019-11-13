package main

import (
	"fmt"
	"time"
)

/*
当启动多个goroutine时，如果其中一个goroutine异常了，
并且我们并没有对进行异常处理，那么整个程序都会终止，
所以我们在编写程序时候最好每个goroutine所运行的函数都做异常处理，异常处理采用recover
*/


func addele(a []int,i int) {
	defer func() {      // 在goroutine 函数内定义一个匿名函数，并使用defer 延迟执行，来捕获异常！（防止主程序直接退出）
		err := recover()
		if err != nil {
			fmt.Println("addele  exec failed")
		}
	}()
	a[i] = i
	fmt.Println("a:",a)
}

func main(){
	arry := make([]int,5)
	for i := 0; i<10; i++ {
		go addele(arry,i)
	}
	time.Sleep(time.Second * 2)

}










