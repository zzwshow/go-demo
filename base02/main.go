package main

import "fmt"

var a string
var b string


func changea(param string) string {
	a += param
	return a
}



func main()  {
	
	// touser := []string{"ym050646","ym050446"}
	// fmt.Println(touser)
	// newtouser := strings.Join(touser,"|")
	// fmt.Println(newtouser)
	if a == ""{
		fmt.Println("ok")
	}
	newa := changea("zzw")
	fmt.Println(newa)
	
}
