package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func f(value *int32, wg *sync.WaitGroup) {
	for i := 0; i < 2000; i++ {
		atomic.AddInt32(value, 1)
	}
	wg.Done()
}

func main() {
	var value int32 = 42
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&value, &wg)
	go f(&value, &wg)
	wg.Wait()

	fmt.Println(value)
}
