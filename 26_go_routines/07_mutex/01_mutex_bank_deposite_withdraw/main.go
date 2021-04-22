package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var balance int

func init() {
	balance = 1000
}
func deposite(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}
func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value

	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go deposite(1000, &wg)
	go withdraw(500, &wg)
	wg.Wait()
	fmt.Printf("New Balance %d\n", balance)

}

/*resource:
https://www.geeksforgeeks.org/mutex-in-golang-with-examples/
https://tour.golang.org/concurrency/9
https://gobyexample.com/mutexes
https://tutorialedge.net/golang/go-mutex-tutorial/
https://golangdocs.com/mutex-in-golang
https://www.golangprograms.com/how-to-use-a-mutex-to-define-critical-sections-of-code-and-fix-race-conditions.html
https://stackoverflow.com/questions/61420524/whats-the-difference-between-channel-and-mutex-in-golang
*/
