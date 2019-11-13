package main

import "fmt"

// 模拟动物行为的接口
type IAnimal interface {
	Eat()  //
}

// 动物所有动物的父类
type Animal struct {
	Name string
}

// 动物去实现IAnimal 接口中的方法
func (a *Animal) Eat(){
	fmt.Printf("%v is eating\n", a.Name)
}

// 动物的构造函数（实例化动物）
func newAnimal(name string) *Animal{
	return &Animal{Name:name}
}

// 定义猫的结构体
type Cat struct {
	*Animal
}

// 实现猫的构造函数 初始化animal结构体
func newCat(name string) *Cat {
	return &Cat{Animal:newAnimal(name)}
}

func (cat *Cat) Eat() {  // 实现cat 的重载，已经不在调用Animal 下的Eat() 方法了
	fmt.Printf("children %v is eating\n", cat.Name)
}


func check(animal IAnimal){  // 定义一个函数，执行接口内的方法
	animal.Eat()
}

func main(){
	cat := newCat("mao")
	var a IAnimal   // 定义一个名为 a 的IAnimal 接口
	a = cat  // 因为cat 结构体对象实现了 IAnimal 接口中的方法 ，所以a接口可以接收cat 结构体对象
	a.Eat()  // 直接调用a 接口的Eat （因为当前a 是cat 结构体对象实现的，所有会执行cat结构体的方法Eat）
	
	check(a) //  将接口传进check 函数，会执行函数内的接口方法
	
	// check(cat)
}

