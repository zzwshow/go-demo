package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type byAge []Person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	family := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(byAge(family))
	fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}
