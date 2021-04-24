package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	// time.Sleep(time.Second)
}

/*An unbuffered channel is a channel that needs a receiver as soon as a message is emitted to the channel.
To declare an unbuffered channel, you just don’t declare a capacity.*/
