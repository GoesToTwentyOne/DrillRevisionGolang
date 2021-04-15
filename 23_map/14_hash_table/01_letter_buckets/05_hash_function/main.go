package main

import "fmt"

func main() {
	v := hash("GO", 12)
	fmt.Println(v)

}
func hash(word string, bucket int) int {
	letter := int(word[0])
	bucket = letter % bucket
	return bucket

}
