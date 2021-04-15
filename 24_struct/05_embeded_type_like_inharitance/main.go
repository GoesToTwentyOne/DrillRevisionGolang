package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}
type secrateAgent struct {
	person
	Id     int
	identy bool
}

func main() {
	s := secrateAgent{
		person: person{
			Fname: "Ahyti",
			Lname: "Byun",
			Age:   22,
		},
		Id:     454545,
		identy: true,
	}
	s1 := secrateAgent{
		person: person{
			Fname: "Alex",
			Lname: "GooT",
			Age:   21,
		},
		Id:     454544444,
		identy: true,
	}
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s.person, s.identy)

}
