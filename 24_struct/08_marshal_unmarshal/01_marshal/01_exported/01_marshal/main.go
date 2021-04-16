package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Fname      string
	Lname      string
	Age        int
	unexported bool
}

func main() {
	p := person{
		Fname:      "Nihad",
		Lname:      "Hossain",
		Age:        21,
		unexported: true,
	}
	p2 := person{
		Fname:      "Neha",
		Lname:      "Jannat",
		Age:        06,
		unexported: true,
	}

	people := []person{p, p2}
	fmt.Println(people) //slice
	fmt.Printf("%#v \n", people)
	bs, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bs)) //json marshal

	//{"Fname":"Nihad","Lname":"Hossain","Age":21},{"Fname":"Neha","Lname":"Jannat","Age":6}
	//json output

}
