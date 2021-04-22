package main

import (
	"fmt"
	"time"
)

var counter int

func main() {
	loop := 0
	for loop < 2 {
		go func() {
			counter++
		}()
		loop++
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Printf("%v \n", counter)

}
