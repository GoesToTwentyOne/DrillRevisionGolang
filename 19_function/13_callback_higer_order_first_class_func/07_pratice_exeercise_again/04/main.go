package main

import "fmt"

func main() {
	goes(func(x []int) {
		fmt.Println(x)
	}, []int{1, 2, 3, 4, 5})

}
func goes(callback func(x []int), numbers []int) {
	// var xs []int
	xs := []int{}

	for _, value := range numbers {
		xs = append(xs, value)

	}
	callback(xs)
}
