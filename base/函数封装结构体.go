package main

import (
	"fmt"
)

type Command struct {
	Name    string    // 指令名称
	Var     *int      // 指令绑定的变量
	Comment string    // 指令的注释
}




func main(){
	var version int = 10
	cmd := newCommand("www",&version,"kkkkk")
	fmt.Println(cmd)
	fmt.Println(cmd.Name)
	fmt.Println(*cmd.Var)
	fmt.Println(cmd.Comment)


}


func newCommand(name string,varref *int,commant string) *Command {
	return &Command{Name:name,Var:varref,Comment:commant}
}