package main

import "fmt"

type heightwidth struct {
	height int
	width  int
}

func (hw heightwidth) trangleArea() float64 {
	return float64(05 * hw.height * hw.width)

}

func main() {
	// var v heightwidth = heightwidth{
	// 	height: 5,
	// 	width:  3,
	// }
	var v heightwidth
	fmt.Println(v)
	v = heightwidth{ //intializition
		height: 5,
		width:  3,
	}
	fmt.Println(v)
	fmt.Printf("%T ", v)
	fmt.Println(v.trangleArea()) //call the method

}
