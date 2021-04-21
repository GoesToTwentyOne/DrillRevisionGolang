package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	started := time.Now()
	var wg sync.WaitGroup

	wg.Add(2)

	go foo(&wg)
	go bar(&wg)

	wg.Wait()
	fmt.Printf("done time %v \n", time.Since(started))

}

func foo(wg *sync.WaitGroup) {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar(wg *sync.WaitGroup) {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}

/*WaitGroup exports 3 methods.

1	Add(int)	 It increases WaitGroup counter by given integer value.
2	Done()	 It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
3	Wait()	It Blocks the execution until it’s internal counter becomes 0.*/
