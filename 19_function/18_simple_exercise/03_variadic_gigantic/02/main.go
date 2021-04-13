package main

import "fmt"

func main() {
	large := gigantic(-20, -30, -700)
	fmt.Println(large)

}
func gigantic(n ...int) int {
	var largest int
	for i, value := range n {
		if value > largest || i == 0 {
			largest = value
		}
	}
	return largest

}

/*What does that code do?
The first time through the range loop
the index, i, will be zero
so largest will be set to the first number
Originally largest is set to the zero value for an int, which is zero
Zero would be greater than any negative number
if you only have negative numbers
you need largest to be something less than zero
Thanks to Ricardo G for this code improvement!
*/
