package main

import "fmt"

func incrementor(s string) <-chan int {
	chOut := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			chOut <- 1
			fmt.Println(s, i)
		}
		close(chOut)
	}()
	return chOut

}

func puller(ch <-chan int) <-chan int {
	chOut := make(chan int)
	go func() {
		var sum int = 0
		for value := range ch {
			sum += value
		}
		chOut <- sum
		close(chOut)

	}()
	return chOut

}

func main() {

	inc := incrementor("goestotwentyone")
	pull := puller(inc)
	inc2 := incrementor("goestotwentyone2")
	pull2 := puller(inc2)
	// for value := range pull {

	// 	fmt.Println(value)
	// }
	fmt.Println("final counter :", <-pull+<-pull2)

}
