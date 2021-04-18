package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 5, 78, 9, 45, 12, 54, 8, 6, 3, 9, 8, 7, 1, 4, 5, 4}
	sort.Ints(s)
	fmt.Println(s)
}
