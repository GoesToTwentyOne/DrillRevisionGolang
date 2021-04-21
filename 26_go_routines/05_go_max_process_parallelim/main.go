package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	/*GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous
	setting. It defaults to the value of runtime.NumCPU. If n < 1, it does not change the current setting.
	 This call will go away when the scheduler improves*/
}

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
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}

/*GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous
setting. It defaults to the value of runtime.NumCPU. If n < 1, it does not change the current setting.
 This call will go away when the scheduler improves*/
