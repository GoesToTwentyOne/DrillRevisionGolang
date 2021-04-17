package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Nihad", "Mahi", "Tonay", "Tahmid", "Neha"}
	sort.StringSlice(s).Sort()
	fmt.Println(s)

}
