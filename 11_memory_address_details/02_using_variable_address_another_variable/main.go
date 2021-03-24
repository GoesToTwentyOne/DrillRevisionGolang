package main

import "fmt"

func main() {
	var a int = 20
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)

}
