package main

import "fmt"

func main() {
	var x [21]int
	fmt.Println(x)
	fmt.Println(len(x))
	x[11] = 88
	fmt.Println(x)
	fmt.Println(x[11])

}

/*Arrays in Golang or Go programming language is much similar to other programming languages.
 In the program, sometimes we need to store a collection of data of the same type, like a list of student marks.
  Such type of collection is stored in a program using an Array. An array is a fixed-length sequence that
  is used to store homogeneous elements in the memory. Due to their fixed length array are not much popular
  like Slice in Go language.
In an array, you are allowed to store zero or more than zero elements in it.
The elements of the array are indexed by using the [] index operator with their zero-based position,
means the index of the first element is array[0] and the index of the last element is array[len(array)-1].*/

//------------------------------------------------------------------------------------------------------------------------------
/*An array is a data structure that consists of a collection of elements of a single type or simply you can
say a special variable, which can hold more than one value at a time. The values an array holds are called
its elements or items. An array holds a specific number of elements, and it cannot grow or shrink.
Different data types can be handled as elements in arrays such as Int, String, Boolean, and others.
The index of the first element of any dimension of an array is 0, the index of the second element of
any array dimension is 1, and so on.*/
