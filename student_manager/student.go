package main

import (
	"fmt"
	"os"
)

type StudemtStruct struct {
	UserName string
	Sex int
	Grade string
	Score float32
}

func NewStudent(username string,sex int,grade string,score float32) (stu *StudemtStruct) {
	stu = &StudemtStruct{
		UserName:username,
		Sex:sex,
		Grade:grade,
		Score:score,
	}
	return stu
}

var AllStudents []*StudemtStruct



func showMenu(){
	fmt.Println("1. add student")
	fmt.Println("2. modify student")
	fmt.Println("3. show all student")
	fmt.Println("4. exited\n\n")
}

func InputStudentInfo() *StudemtStruct{
	var (
		username string
		sex int
		grade string
		score float32
	)
	fmt.Println("please input username: ")
	fmt.Scanf("%s\n",&username)
	fmt.Println("please input  sex:[0|1]: ")
	fmt.Scanf("%d\n",&sex)
	fmt.Println("please input  grade:[0-6]: ")
	fmt.Scanf("%s\n",&grade)
	fmt.Println("please input  score:[0-100]: ")
	fmt.Scanf("%f\n",&score)

	stu := NewStudent(username,sex,grade,score)
	return stu
}


func AddStudent(){
	stu := InputStudentInfo()
	for index,v := range AllStudents{  //去重
		if(v.UserName == stu.UserName){
			AllStudents[index] = stu
			return
		}
	}
	AllStudents = append(AllStudents,stu)
	fmt.Printf("user: %s add success\n",stu.UserName)


}




func ModifyStudent(){
	stu := InputStudentInfo()
	for index,v := range AllStudents{  //去重
		if(v.UserName == stu.UserName){
			AllStudents[index] = stu
			fmt.Printf("user: %s Modify success\n",stu.UserName)
			return
		}
	}
	fmt.Printf("user: %s is not found\n",stu.UserName)

}

func ShowAllStudent(){
	for _,v := range AllStudents{
		fmt.Printf("user %s info:%#v\n",v.UserName,v)
	}
}




func main() {
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n",&sel)
		switch sel {
		case 1:
			AddStudent()
		case 2:
			ModifyStudent()
		case 3:
			ShowAllStudent()
		case 4:
			os.Exit(0)




		}
	}

	
}
