package main

import (
	"fmt"
	"reflect"
)

// 反射:
// 类型太多,类型断言猜不全,使用反射就能直接拿到接口值的动态类型和动态值

// typeOf()

func reflectType(x interface{}){
	v:= reflect.TypeOf(x)
	fmt.Printf("Type:%v\n",v)

	// 拿到类型和种类
	fmt.Printf("type:%v, kind:%v",v.Name(),v.Kind()) //type:Cat, kind:struct

}

type Cat struct {
	name string
}


func main(){
	reflectType(100) // Type:int
	reflectType([3]int{1,2,3}) // Type:[3]int
	reflectType(map[string]int{}) // Type:map[string]int


	// 测试自定义的类型  在go 中可以造任意类型 所以分为,类型(type)和种类(kind)
	var c1 = Cat{name:"花猫"} // Type:main.Cat
	reflectType(c1)

	var age int = 20
	reflectType(&age) //type:, kind:ptr 指针类型为空

}
