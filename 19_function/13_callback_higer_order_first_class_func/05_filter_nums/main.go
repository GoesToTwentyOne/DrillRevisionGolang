package main

import "fmt"

func main() {
	final := goes(func(value int) bool {
		return value > 1

	}, []int{1, 2, 3})
	fmt.Println(final)

}
func goes(callback func(int) bool, numbers []int) []int {
	var xs []int
	for _, value := range numbers {
		if callback(value) {
			xs = append(xs, value)
		}

	}
	return xs
}
