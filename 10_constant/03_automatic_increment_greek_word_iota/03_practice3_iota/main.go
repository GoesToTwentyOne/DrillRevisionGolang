package main

import "fmt"

const (
	a = iota
	b
	c
)
const (
	e = iota
	f
	g
)

func main() {
	fmt.Println("           ", a, b, c, e, f, g)

}
