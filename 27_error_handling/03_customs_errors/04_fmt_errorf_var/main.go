package main

import (
	"fmt"
	"log"
	"math"
)

func main() {

	value, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		errorNew := fmt.Errorf("norgate math: square root of negative number %v", n)
		return 0, errorNew
	}
	return math.Sqrt(n), nil
}
