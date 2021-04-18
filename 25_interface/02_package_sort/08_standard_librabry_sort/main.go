package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("HI aaaaaaa    %s   %d", p.Name, p.Age)
}

type age []person

func (a age) Len() int {
	return len(a)

}
func (a age) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]

}
func (a age) Less(i, j int) bool {
	return a[i].Age < a[j].Age

}

func main() {
	p := []person{
		{"Noq", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(p[0])
	fmt.Println(p)
	sort.Sort(age(p))
	fmt.Println(p)
}
