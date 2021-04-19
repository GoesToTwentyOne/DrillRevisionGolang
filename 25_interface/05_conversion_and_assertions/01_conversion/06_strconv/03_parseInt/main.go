package main

import (
	"fmt"
	"strconv"
)

func main() {
	//ParseInt,ParseFloat64,ParseBool,ParseUint   are converts string to type of value
	fmt.Println("-------------------------  Int64----------------------------------")
	a, err := strconv.ParseInt("-400", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a)
	fmt.Printf("%T   \n", a)
	fmt.Println("-------------------------  Float64----------------------------------")
	b, err := strconv.ParseFloat("400", 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(b)
	fmt.Printf("%T   \n", b)
	fmt.Println("-------------------------  UnInt----------------------------------")
	c, err := strconv.ParseUint("100", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c)
	fmt.Printf("%T   \n", c)
	fmt.Println("-------------------------  Bool----------------------------------")
	d, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)
	fmt.Printf("%T   \n", d)

}
