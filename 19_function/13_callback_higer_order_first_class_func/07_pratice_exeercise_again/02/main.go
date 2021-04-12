package main

import "fmt"

func main() {
	fmt.Println(goes(func(x []int) int {
		var sum int = 0
		for _, value := range x {
			sum += value

		}
		return sum

	}, []int{1, 2, 3, 4, 5, 6}))

}
func goes(f func(x []int) int, n []int) int {
	odd := f(n)
	return odd

}
