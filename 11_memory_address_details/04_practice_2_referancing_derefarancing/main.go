package main

import "fmt"

func main() {
	y := 4
	fmt.Println("before ", y)
	goes(&y)
	fmt.Println("after", y)

}
func goes(x *int) {
	fmt.Printf("before %v  \n", x)
	fmt.Printf("before %v \n", *x)
	*x = 50
	fmt.Printf(" %v after \n", x)
	fmt.Printf(" %v after \n", *x)

}
