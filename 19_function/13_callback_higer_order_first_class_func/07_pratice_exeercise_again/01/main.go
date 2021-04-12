package main

import "fmt"

func main() {
	v := func(xi []int) int {
		if len(xi) == 0 {
			return 0

		}
		if len(xi) == 1 {
			return 1
		}
		return xi[0] + xi[len(xi)-1]

	}
	final := goes(v, []int{1, 2, 3, 4})
	fmt.Println(final)

}
func goes(f func(xi []int) int, ii []int) int {
	n := f(ii)
	return n

}
