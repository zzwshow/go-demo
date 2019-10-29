package main

import (
	"encoding/json"
	"fmt"
)


func main()  {
	m := make(map[string]interface{},4)
	m["A"] = "abc"
	m["B"] = []string{"A","B","C"}
	m["C"] = true
	m["D"] = 123

	//编码成json
	//result,err := json.Marshal(m)
	result ,err := json.MarshalIndent(m,"" ," ")
	if err != nil{
		fmt.Println("err =",err)
		return
	}
	fmt.Println("result = ", string(result))

}
