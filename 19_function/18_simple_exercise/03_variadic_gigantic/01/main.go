package main

import "fmt"

func main() {
	great := gigantic(1000, 1, 50, 87, 101, 500)
	fmt.Println(great)

}
func gigantic(n ...int) int {
	var largest int
	for _, value := range n {
		if value > largest {
			largest = value
		}
	}
	return largest

}
