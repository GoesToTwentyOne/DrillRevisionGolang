package main

import "fmt"

func increment() chan int {
	chOut := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			chOut <- i
		}
		close(chOut)

	}()
	return chOut

}
func sumInc(s chan int) chan int {
	chOut := make(chan int)
	var sum int = 0
	go func() {
		for values := range s {
			sum += values
		}
		chOut <- sum
		close(chOut)

	}()
	return chOut

}

func main() {
	chpass := increment()
	result := sumInc(chpass)
	// fmt.Println(<-result)
	for r := range result {
		fmt.Println(r)
	}

}
