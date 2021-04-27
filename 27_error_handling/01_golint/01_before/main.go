package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	} else {
		return fmt.Sprint("x is less than 10")
	}
}

/*Golint is a useful Go linting tool that prints out coding style mistakes.
While it's not perfect, it presents suggestions that follow many of the items in Effective Go and
the CodeReviewComments wiki.*/
