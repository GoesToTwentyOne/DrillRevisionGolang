package main

import "fmt"

func receiveChannel(ch chan int) {
	//data received
	fmt.Println(235 + <-ch)
}

func main() {
	ch := make(chan int)
	go receiveChannel(ch)

	//data send
	ch <- 25

}
