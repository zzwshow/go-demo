package main

import "fmt"

//


type person struct {  // 声明一个新的类型
	name string
	age int
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age{
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}


func main(){
	var tom person  // 定义一个tom 变量 类型是person
	tom.name,tom.age = "tom",28  //赋值初始化
	
	bob := person{name:"Bob",age:25}  //定义变量 并指明字段赋值实例化
	paul := person{"paul",43} //定义变量 按照结构体属性定义的顺序赋值
	
	tb_Older,tb_diff := Older(tom,bob)
	tp_Older,tp_diff := Older(tom,paul)
	
	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)
	
	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)
	


}


