package main

import "fmt"

const (
	_a = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf(" %b \t", kb)
	fmt.Printf(" %d \t", kb)
	fmt.Printf("%b \t", mb)
	fmt.Printf(" %d \t", mb)
	fmt.Printf("%b \t", gb)
	fmt.Printf(" %d \t", gb)

}
