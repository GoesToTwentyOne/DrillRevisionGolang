package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	a := person{Fname: "Nihad", Lname: "Hossain", Age: 21}
	b := person{Fname: "Rowjatul", Lname: "Neha", Age: 06}
	fmt.Println(a.Fname, a.Lname, a.Age)
	fmt.Println(b.Fname, b.Lname, b.Age)
	fmt.Println(a)
	fmt.Println(b)

}
