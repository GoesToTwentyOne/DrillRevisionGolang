package main

import (
	"fmt"
	"sync"
)

var (
	wg      sync.WaitGroup
	counter int
	mutex   sync.Mutex
)

func incrementor(s string) {
	mutex.Lock()
	for i := 0; i < 20; i++ {
		counter++
		fmt.Println(s, i, "counter : ", counter)
	}
	mutex.Unlock()
	wg.Done()

}
func main() {
	wg.Add(2)
	go incrementor("Goes :")
	go incrementor("Goes 2 :")
	wg.Wait()
	fmt.Println("Final Counter:", counter)

}
