package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	value, err := Sqrt(10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(value)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	// implementation  //default that
	return math.Sqrt(f), nil
}
