package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break
		}
	}

}

// When a break statement is encountered inside a loop,
// the loop is immediately terminated and the program control resumes at the next statement following the loop.
