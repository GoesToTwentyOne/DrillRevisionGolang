package main

import "fmt"

func main() {
	var x []string
	var y [][]string
	// x[0] = "Nihad" panic error
	x = append(x, "Nihad")
	fmt.Println(x)
	fmt.Println(y)
}
