package main

import "fmt"

func channelClose(ch chan string) {
	for i := 0; i < 5; i++ {
		//data send
		ch <- "golang simplicity"
	}
	close(ch)
}

func main() {
	//create channel

	// var ch chan string
	ch := make(chan string)

	go channelClose(ch)

	//data receuve

	for {
		res, ok := <-ch
		//if ok == false
		if !ok {
			fmt.Printf("channel Close %v %v: \n", res, ok)
			break
		}
		fmt.Printf("channel open : %v %v\n", res, ok)
	}

}
