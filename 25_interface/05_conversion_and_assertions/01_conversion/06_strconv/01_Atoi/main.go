package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var x = "15"
	var y = 6
	z, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(y + z)
}
