package main

import "fmt"

func main() {
	fact := factorial(4)
	for value := range fact {
		fmt.Println(value)
	}

}

func factorial(n int) chan int {
	ch := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i

		}
		ch <- total
		close(ch)
	}()
	return ch

}
