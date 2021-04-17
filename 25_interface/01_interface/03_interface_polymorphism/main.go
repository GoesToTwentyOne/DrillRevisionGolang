package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	height int
	weight int
}
type circle struct {
	radius int
}

func (r rectangle) area() float64 {
	return (float64(r.height * r.weight))

}
func (c circle) area() float64 {
	return float64(c.radius*c.radius) * math.Pi

}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())

}

func main() {
	finalRectangle := rectangle{
		height: 4,
		weight: 5,
	}
	finalCircle := circle{
		radius: 5,
	}
	info(finalRectangle) //{ Polymorphism}
	info(finalCircle)

}
