package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("---------------------------------------------------------------------------")
	for i := 0; i < 90; i++ {
		x = append(x, i)
		fmt.Println("Length :", len(x), "-----", "Value : ", x[i], "----", "Capacity: ", cap(x))
	}

}
