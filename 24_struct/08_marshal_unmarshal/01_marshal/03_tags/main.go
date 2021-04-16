package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string `json:"Naming first"`
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := person{"James", "Bond", 20}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bs))
}

//https://golang.org/pkg/encoding/json/
