package main

import "fmt"

func main() {
	final := even(func(n int) bool {
		return n%2 == 0
	}, []int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(final)

}

func even(callback func(n int) bool, numbers []int) []int {
	var odd []int
	for _, value := range numbers {
		if callback(value) {
			odd = append(odd, value)
		}
	}
	return odd

}
