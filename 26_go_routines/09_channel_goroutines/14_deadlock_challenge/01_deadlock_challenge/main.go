package main

import "fmt"

func main() {

	ch := make(chan int)
	ch <- 45
	fmt.Println(<-ch)
}
