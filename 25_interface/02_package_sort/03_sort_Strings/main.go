package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Nihad", "Mahi", "Tonay", "Tahmid", "Neha"}
	sort.Strings(s)
	fmt.Println(s)
}
