package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		fmt.Println("channel 1 ", <-done)
		fmt.Println("channel 2 ", <-done)

		// <-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
