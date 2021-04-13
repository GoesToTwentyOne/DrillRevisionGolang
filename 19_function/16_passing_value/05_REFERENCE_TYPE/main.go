package main

import "fmt"

func main() {
	// v := new(map[string]int)      //error
	// chaneMe(v)

	v := make(map[string]int)
	changeME(v)
	fmt.Println(v["Nihad"])

}
func changeME(n map[string]int) {
	n["Nihad"] = 21
}
