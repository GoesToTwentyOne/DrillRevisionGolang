package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "Nihad",
		1: "Alex",
		2: "Anna",
		3: "Candy",
	}
	fmt.Println(myMap)
	// delete(myMap, 2)
	if val, exist := myMap[2]; exist {
		fmt.Println("This map is exist")
		fmt.Println("value: ", val)
		fmt.Println("Exist: ", exist)
	} else {
		fmt.Println("This map isn't exist")
		fmt.Println("value: ", val)
		fmt.Println("Exist: ", exist)
	}

}
