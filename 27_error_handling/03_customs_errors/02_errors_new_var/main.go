package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var errorNew = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T \n", errorNew)
	value, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errorNew
	}
	return math.Sqrt(n), nil
}
