package main

import "fmt"

func main() {
	oddarg := oddCover()
	fmt.Println(oddarg())
	fmt.Println(oddarg())
	fmt.Println(oddarg())
	fmt.Println(oddarg())
	//another object
	fmt.Println("another object")

	oddarg2 := oddCover()
	fmt.Println(oddarg2())
	fmt.Println(oddarg2())
	fmt.Println(oddarg2())
	fmt.Printf("%T \n", oddarg)

}
func oddCover() func() int {
	i := 0
	return func() int {
		i += 2
		return i

	}

}
