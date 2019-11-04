package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	jsonData := `
	{
	"company": "xxx",
	"subjects": ["GO","PYTHON","JAVA"],
	"isok": true,
	"price":66.66
	}`

	m := make(map[string]interface{},4)

	err := json.Unmarshal([]byte(jsonData),&m)
	if err != nil{
		fmt.Println("err = ",err)
	}
	fmt.Println("m = ",m)  //现在对map没法直接使用

	//类型断言
	for key, value := range m {
		fmt.Printf("%v------------%v\n",key,value)
	}

}