package main

import "fmt"

func main() {
	x := 21
	fmt.Println(x)
	{
		//access in clouser x
		fmt.Println(x)
		name := "Nihad Hossain"
		fmt.Println(name)
	}
	//don't use the out of clouser
	//fmt.Println(name)
}
