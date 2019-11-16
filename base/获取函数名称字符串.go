package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// 定义一个测试函数
func foo(){

}


// 定义一个返回函数名称字符串的方法
func GetFuncNmae(i interface{},seps ...rune) string {  // 接收一个任意类型的函数名，和任意多个字符用来切割
	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	// fmt.Println(fn)  //main.foo
	
	// 用seps 进行切割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _,s := range seps {   //这个seps 是我们需要切割掉的字符，
			if sep == s {
				return true       //如歌fn 中包含定义的字符就切掉
			}
		}
		return false  // 找完之后返回 数组
	})
	//  去除字符后 ： [main foo]
	fmt.Println(fields[1])  //foo
	return fn
}



func main(){
	GetFuncNmae(foo,'.')


}
