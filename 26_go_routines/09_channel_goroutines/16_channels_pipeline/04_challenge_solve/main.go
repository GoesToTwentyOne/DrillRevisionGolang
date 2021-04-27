package main

import "fmt"

func main() {
	channel := gen()
	finalchannel := sq(channel)
	for final := range finalchannel {
		fmt.Println(final)
	}

}

func gen() <-chan int {
	chOut := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				chOut <- j
			}
		}
		close(chOut)
	}()
	return chOut
}
func sq(input <-chan int) <-chan int {
	chOut := make(chan int)
	go func() {
		for value := range input {
			chOut <- fact(value)
		}
		close(chOut)
	}()
	return chOut

}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
