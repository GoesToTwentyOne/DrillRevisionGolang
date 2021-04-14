package main

import "fmt"

func main() {
	x := make([]string, 3, 5)
	//length-number of the elements referred to by the slice
	//capacity-number of the elements  in the undeclaring array
	x[0] = "Nihad"
	x[1] = "Roja"
	x[2] = "Shova"
	x = append(x, "Alex")
	x = append(x, "Juse")
	x = append(x, "Anna")
	x = append(x, "Emma")
	fmt.Println(x[6])
	fmt.Println(x)

}
