package main

import "fmt"

func main() {
	ch := make(chan string, 10)
	go func() {
		ch <- "Golang simplicity"
		ch <- "Golang concurrency"
		ch <- "Golang statically typed"
		ch <- "Golang garbage collective"
		ch <- "Golang super fast"
		close(ch)

	}()
	for resValue := range ch {
		fmt.Println(resValue)

	}
	// fmt.Println("Length : ", len(ch))
	fmt.Println("Capacity : ", cap(ch))

}
