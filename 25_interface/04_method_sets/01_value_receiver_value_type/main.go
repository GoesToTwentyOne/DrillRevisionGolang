package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return (math.Pi * c.radius * c.radius)

}
func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	c := circle{
		radius: 5,
	}
	info(c)

}
