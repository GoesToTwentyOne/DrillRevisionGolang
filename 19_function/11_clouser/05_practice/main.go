package main

import "fmt"

func wra() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	inc1 := wra()
	inc2 := wra()
	fmt.Println("1:", inc1())
	fmt.Println("1:", inc1())
	fmt.Println("2:", inc1())
	fmt.Println("2:", inc2())
	fmt.Println("2:", inc2())
}
