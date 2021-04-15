package main

import "fmt"

func main() {
	var myMap map[string]string
	fmt.Println(myMap)
	fmt.Println(myMap == nil)

}

/*In Go language, a map is a powerful, ingenious, and a versatile data structure.
Golang Maps is a collection of unordered pairs of key-value. It is widely used because
it provides fast lookups and values that can retrieve, update or delete with the help of keys.*/
//The make function returns a map of the given type, initialized and ready for use.

/* if add value like this myMap["A"]="Absulattely"  return an errror*/
