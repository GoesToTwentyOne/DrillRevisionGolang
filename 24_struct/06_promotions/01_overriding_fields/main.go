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

func main() {
	s := student{
		Department: "CSE",
		Roll:       399999999999,
		Major:      "Software Engineering",
		person: person{
			Fname: "Nihad",
			Lname: "Hossain",
			Age:   21,
		},
	}
	s2 := student{
		Department: "MBBS",
		Roll:       45555555555,
		Major:      "Medicine",
		person: person{
			Fname: "Sanzida",
			Lname: "Shova",
			Age:   21,
		},
	}
	// Promoted fields of the student structure
	fmt.Println(s.Fname, s.Lname, s.Age)
	fmt.Println(s2.Fname, s2.Lname, s2.Age)

	//normal field of student structue
	fmt.Println(s.Department, s.Major, s.Roll)
	fmt.Println(s2.Department, s2.Major, s2.Roll)

}

//https://www.geeksforgeeks.org/promoted-fields-in-golang-structure/
/*In Go structure, promoted fields are just like anonymous fields, the type of the field is the name of the
field. We use this concept in the nested structure where a structure is a field in another structure,
simply by just adding the name of the structure into another structure and it behaves like the Anonymous
Field to the nested structure. And the fields of that structure (other than nested structure) are the part
 of the nested structure, such type of fields are known as Promoted fields. If the anonymous structure or
 nested structure and parent structure contains a field that has the same name, then that field
 doesn’t promote, only different name fields get promoted to the structure.*/
