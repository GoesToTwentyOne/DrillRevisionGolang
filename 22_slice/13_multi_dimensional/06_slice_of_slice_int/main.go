package main

import "fmt"

func main() {
	x := make([][]int, 0, 5)
	for i := 0; i < 5; i++ {
		y := make([]int, 0, 7)
		for j := 0; j < 7; j++ {
			y = append(y, j)
		}
		x = append(x, y)
	}
	fmt.Println(x)

}
