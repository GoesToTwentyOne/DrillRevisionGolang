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
	fmt.Println(string(bs))

}
