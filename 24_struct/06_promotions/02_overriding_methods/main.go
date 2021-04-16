package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}
type student struct {
	Department string
	Roll       int
	Major      string
	person
}

func (p person) greeting() {
	fmt.Println("I'm just a regular person")
}
func (p student) greeting() {
	fmt.Println("I'm just a regular student")
}

func main() {
	s := student{
		Department: "CSE",
		Roll:       455555,
		Major:      "Software Engineering",
		person: person{
			Fname: "Nihad",
			Lname: "Hossain",
			Age:   21,
		},
	}
	s2 := student{
		Department: "MBBS",
		Roll:       777788,
		Major:      "Medicine",
		person: person{
			Fname: "Sanzida",
			Lname: "Shova",
			Age:   21,
		},
	}
	s.greeting()
	s2.greeting()
	s.person.greeting()

}
