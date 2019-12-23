package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {

	var persions []*Person

	for i := 0; i < 10; i++ {
		p := &Person{
			Name: fmt.Sprintf("test %d", i),
			Age:  29,
			Sex:  "男"}

		persions = append(persions, p)
	}

	data, err := json.Marshal(persions)
	if err != nil {
		fmt.Println("= marshal failed,err %v", err)
		return
	}

	err = ioutil.WriteFile("./demo.txt", data, 0755)
	if err != nil {
		fmt.Printf("write file failed err:%v", err)
		return
	}
}
