package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk bool	`json:"isok"`
	Price float64 `json:"price"`
}

type IT2 struct {
	Subjects []string `json:"subjects"`
}

func main()  {

	jsonData := `
	{
	"company": "xxx",
	"subjects": ["GO","PYTHON","JAVA"],
	"isok": true,
	"price":66.66
}`

	var tmp IT // 定义一个结构体变量
	err := json.Unmarshal([]byte(jsonData),&tmp)

	if err != nil{
		fmt.Println("err = ",err)
		return
	}
	fmt.Printf("%T\n",tmp) //类型已经转换为结构体了    main.IT
	fmt.Println("tmp = ", tmp)  //tmp =  {xxx [GO PYTHON JAVA] true 66.66}
	fmt.Println("Price:",tmp.Price) //Price: 66.66

	// 只解析我们需要的字段赋值给结构体
	var tmp2 IT2
	err2 := json.Unmarshal([]byte(jsonData),&tmp2)

	if err2 != nil{
		fmt.Println("err2 = ",err2)
		return
	}
	fmt.Printf("%T\n",tmp2) //类型已经转换为结构体了
	fmt.Println("tmp2 = ", tmp2) //tmp2 =  {[GO PYTHON JAVA]}


}