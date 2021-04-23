package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

func f(v *int, wg *sync.WaitGroup) {
	for i := 0; i < 3000; i++ {
		*v++
	}
	wg.Done()
}

func main() {
	var v int = 42
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&v, &wg)
	go f(&v, &wg)
	wg.Wait()

	fmt.Println(v)
}

//go run -race main.go

/*
https://golangdocs.com/atomic-operations-in-golang-atomic-package
https://golang.org/pkg/sync/atomic/
https://www.geeksforgeeks.org/atomic-variable-in-golang/
https://www.golangprograms.com/how-to-fix-race-condition-using-atomic-functions-in-golang.html
*/
