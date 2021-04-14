package main

import "fmt"

func main() {
	x := make([]int, 4)
	// 4 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	x[0] = 4
	x[1] = 5
	x[2] = 6
	x[3] = 7
	fmt.Println(x)
	fmt.Println(len(x)) //length an array
	fmt.Println(cap(x)) //capacity an arra
	y := make([]int, 4, 6)
	fmt.Println("---------------------------------------------------------------------------------------")
	// 4 is length - number of elements referred to by the slice
	// 6 is capacity - number of elements in the underlying array
	// you could also do it like this
	y[0] = 4
	y[1] = 5
	y[2] = 6
	y[3] = 7
	fmt.Println(len(y))
	fmt.Println(cap(y))
	fmt.Println(y)

}
