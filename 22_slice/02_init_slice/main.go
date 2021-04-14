package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15, 14}
	for i, value := range x {
		fmt.Printf("%v---------%v  \n", i, value)
	}

}
