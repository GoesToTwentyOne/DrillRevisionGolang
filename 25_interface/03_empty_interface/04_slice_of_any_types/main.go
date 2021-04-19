package main

import "fmt"

type animal struct {
	sound string
}
type dog struct {
	animal
	friendly bool
}
type cat struct {
	animal
	annoying bool
}

func main() {
	d := dog{
		animal: animal{
			sound: "buh buh",
		},
		friendly: true,
	}
	c := cat{
		animal: animal{
			sound: "Maw maw",
		},
		annoying: true,
	}
	critters := []interface{}{d, c}
	fmt.Println(critters)

}
