package main

import "fmt"

func main() {
	i := inc()
	fmt.Println(i())
	fmt.Println(i())

}
func inc() func() int {
	var x int
	return func() int {
		x++
		return x

	}

}
