package main

import "fmt"

func main() {
	mul(func(a, b, c int) int {
		return a * b * c

	})

}
func mul(i func(a, b, c int) int) {
	fmt.Println(i(10, 20, 30))
}
