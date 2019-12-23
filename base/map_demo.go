package main

import "fmt"

func main() {
	test := map[string]interface{}{"a": 1, "b": "b", "release_time": "2019-12-21"}
	_, queryTime := test["release_time"]

	delete(test, "release_time")

	for _, v := range test {
		fmt.Println(v)
	}

	fmt.Println(queryTime)
}
