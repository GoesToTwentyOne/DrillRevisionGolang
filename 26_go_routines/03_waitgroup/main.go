package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	started := time.Now()

	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
	fmt.Printf("done time %v \n", time.Since(started))

}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}

/*WaitGroup exports 3 methods.

1	Add(int)	 It increases WaitGroup counter by given integer value.
2	Done()	 It decreases WaitGroup counter by 1, we will use it to indicate termination of a goroutine.
3	Wait()	It Blocks the execution until it’s internal counter becomes 0.*/
