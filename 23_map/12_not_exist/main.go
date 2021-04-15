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
	if value, exist := myMap[7]; exist {
		fmt.Println("exist the code")
		fmt.Println("value : ", value)
		fmt.Println("exist : ", exist)
	} else {
		fmt.Println("don't exist the code")
		fmt.Println("value : ", value)
		fmt.Println("exist : ", exist)

	}
	fmt.Println(myMap)

}
