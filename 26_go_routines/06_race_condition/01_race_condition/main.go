package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

/*Race conditions occur due to unsynchronized access to shared resource and attempt to read and write to
that resource at the same time.Atomic functions provide low-level locking mechanisms for synchronizing
access to integers and pointers. Atomic functions generally used to fix the race condition.
The functions in the atomic under sync packages provides support to synchronize goroutines by locking
access to shared resources.*/
