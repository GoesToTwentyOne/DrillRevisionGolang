package main

import "fmt"

func main() {
	y := 5
	goes(&y)
	fmt.Println(y)

}
func goes(x *int) {
	*x = 2

}
