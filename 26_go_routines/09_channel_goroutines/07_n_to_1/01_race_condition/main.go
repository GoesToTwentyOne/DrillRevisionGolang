package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			wg.Add(1)
			ch <- i
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()

}
