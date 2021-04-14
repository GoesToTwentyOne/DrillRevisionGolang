package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:2], y[3:]...)
	fmt.Println(x)

}
