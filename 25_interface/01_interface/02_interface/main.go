package main

import "fmt"

type heightwidth struct {
	height int
	width  int
}

func (hw heightwidth) trangleArea() float64 {
	return float64(hw.height * hw.width)

}

type shape interface {
	trangleArea() float64
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.trangleArea())
}

func main() {
	final := heightwidth{
		height: 5,
		width:  4,
	}
	fmt.Printf(" %T  \n", final)
	info(final)

}
