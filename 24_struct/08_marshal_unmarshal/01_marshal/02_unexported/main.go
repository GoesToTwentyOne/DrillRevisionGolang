package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	fname      string
	lname      string
	age        int
	unexported bool
}

func main() {
	p := person{
		fname:      "Nihad",
		lname:      "Hossain",
		age:        21,
		unexported: true,
	}
	// p2 := person{
	// 	Fname:      "Neha",
	// 	Lname:      "Jannat",
	// 	Age:        06,
	// 	unexported: true,
	// }
	bs, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bs)) //{}

}
