package main

import "fmt"

type parallelogram struct {
	base   float64
	height float64
}
type trapezoid struct {
	width1 float64
	width2 float64
	height float64
}

func (p parallelogram) area() float64 {
	return (p.base * p.height)

}
func (t trapezoid) area() float64 {
	return (.5 * (t.width1 + t.width2) * t.height)

}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())

}
func totalArea(s ...shape) float64 {
	var sum float64 = 0
	for _, value := range s {
		sum += value.area()
	}

	return sum
}

func main() {
	final1 := parallelogram{
		height: 5,
		base:   4,
	}
	final2 := trapezoid{
		width1: 4,
		width2: 5,
		height: 4,
	}
	info(final1)
	info(final2)
	tt := totalArea(final1, final2)
	fmt.Printf("Total : %v\n", tt)

}
