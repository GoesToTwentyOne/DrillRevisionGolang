package main

import (
	"fmt"
	"strconv"
)

func main() {
	//FormatInt,FormatFloat64,FormatBool,FormatUint   are converts value to string
	fmt.Println("-------------------------  string ----------------------------------")
	a := strconv.FormatInt(-400, 10)
	fmt.Println(a)
	fmt.Printf("%T   \n", a)
	fmt.Println("-------------------------  string ----------------------------------")
	b := strconv.FormatFloat(45.888, 'E', -1, 64)
	fmt.Println(b)
	fmt.Printf("%T   \n", b)
	fmt.Println("-------------------------  string----------------------------------")
	c := strconv.FormatUint(100, 10)
	fmt.Println(c)
	fmt.Printf("%T   \n", c)
	fmt.Println("-------------------------  string----------------------------------")
	d := strconv.FormatBool(true)
	fmt.Println(d)
	fmt.Printf("%T   \n", d)

}
