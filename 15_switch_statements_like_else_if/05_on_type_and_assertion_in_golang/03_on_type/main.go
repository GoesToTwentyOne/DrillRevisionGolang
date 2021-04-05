package main

import "fmt"

func main() {
	taking(7)
	taking("Nihad Hossain")
	var t = contact{"Nihad Hossain", "Hello"}
	taking(t)
	taking(t.Name)
	taking(t.Greeting)

}

type contact struct {
	Name     string
	Greeting string
}

func taking(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("this is integer")
	case string:
		fmt.Println("this is string")
	case contact:
		fmt.Println("This is contact type")
	default:
		fmt.Println("unknown something")
	}
}
