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

func shape(a interface{}) {
	fmt.Println(a)
}

func main() {
	d := dog{
		animal: animal{
			sound: "Good",
		},
		friendly: true,
	}
	c := cat{
		animal: animal{
			sound: "Good",
		},
		annoying: true,
	}
	shape(d)
	shape(c)

}
