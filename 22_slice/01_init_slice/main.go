package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("%T \n", x)
	fmt.Println(x)
	fmt.Println(x[:2])

}

/*In Go language slice is more powerful, flexible, convenient than an array, and is a lightweight data structure.
 Slice is a variable-length sequence which stores elements of a similar type, you are not allowed to store
 different type of elements in the same slice. It is just like an array having an index value and length,
 but the size of the slice is resized they are not in fixed-size just like an array. Internally, slice and
 an array are connected with each other, a slice is a reference to an underlying array. It is allowed to
store duplicate elements in the slice.
 The first index position in a slice is always 0 and the last one will be (length of slice – 1).*/
