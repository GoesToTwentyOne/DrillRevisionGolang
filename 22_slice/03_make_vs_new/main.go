package main

import "fmt"

func main() {
	var x = new(int)
	fmt.Println(x)
	fmt.Println(*x)

	var y = new([]int)
	fmt.Println(y)
	fmt.Println(*y)

	var z = make([]int, 10, 100)
	fmt.Println(z)
}
