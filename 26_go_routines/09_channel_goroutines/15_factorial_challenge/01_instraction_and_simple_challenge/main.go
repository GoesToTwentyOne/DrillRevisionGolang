package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println("factorial : ", f)

}
func factorial(num int) int {
	fact := 1
	for i := num; i > 0; i-- {
		fact *= i

	}
	return fact

}
