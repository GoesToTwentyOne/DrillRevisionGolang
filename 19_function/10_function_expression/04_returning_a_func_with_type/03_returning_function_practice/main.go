package main

import "fmt"

func main() {
	m := mul()
	m1 := m(4.5, 3.5)
	fmt.Println(m1)

}
func mul() func(x, y float64) float64 {
	return func(x, y float64) float64 {
		return x * y
	}

}

//https://www.ardanlabs.com/blog/2013/10/functions-and-naked-returns-in-go.html
