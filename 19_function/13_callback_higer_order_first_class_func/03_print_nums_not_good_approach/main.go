package main

import "fmt"

func main() {
	v := func(n []int) int {
		var intnumber int
		for _, value := range n {
			intnumber, _ = fmt.Println(value)

		}
		return intnumber

	}
	final := show(v, []int{1, 2, 3, 4})
	fmt.Println(final)

}
func show(f func(n []int) int, numbers []int) int {

	g := f(numbers)
	return g
}
