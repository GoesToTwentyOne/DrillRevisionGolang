package main

import "fmt"

type security struct {
	Name       string
	Age        int
	IdentityNo int
	Type       string
}

func main() {
	v := security{Name: "Nihad", Age: 21, IdentityNo: 4587, Type: "Central Security Forces"}
	goes(&v)

}
func goes(s *security) {
	fmt.Println(s.Name)
	fmt.Println(&s.Name)
	s.Name = "Neha"
	fmt.Println(s.Name)
	fmt.Println(&s.Name)

}
