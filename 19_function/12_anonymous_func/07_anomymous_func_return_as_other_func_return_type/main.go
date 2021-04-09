package main

import "fmt"

func main() {
	v := gph()
	fmt.Println(v(20, 30, 40))

}
func gph() func(a, b, c int) int {
	b := func(a, b, c int) int {
		return a * b * c
	}
	return b
}
