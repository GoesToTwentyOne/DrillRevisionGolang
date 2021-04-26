package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()

	for values := range ch {
		fmt.Println(values)
	}
}

// remember to close your channel
// if you do not close your channel, you will receive this error
// fatal error: all goroutines are asleep - deadlock!

// ************** IMPORTANT **************
// YOU NEED GO VERSION 1.5.2 OR GREATER
// otherwise you will receive this error
// fatal error: all goroutines are asleep - deadlock!
