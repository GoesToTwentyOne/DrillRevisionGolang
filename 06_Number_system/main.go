package main

import "fmt"

func main() {
	//%d-> decemal
	//%b -> binary
	// %x -> hexadecimal
	// \t -> tab scabe character
	// %q -> quation with hexadecimal
	fmt.Printf("%d  %b  %x \n", 21, 21, 21)
	fmt.Printf("%d  \t %b  \t %#x \t %q  \n", 21, 21, 21, 21)

}
