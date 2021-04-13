package main

import "fmt"

type customer struct {
	Name string
	Age  int
}

func main() {
	c := customer{Name: "Nihad", Age: 44}
	fmt.Println(c.Name)
	fmt.Println(&c.Name)
	fmt.Println(c.Age)
	fmt.Println(&c.Age)
	changeME(&c)
	fmt.Println(c.Name)
	fmt.Println(&c.Name)
	fmt.Println(c.Age)
	fmt.Println(&c.Age)

}
func changeME(x *customer) {
	fmt.Println(&x.Name)
	fmt.Println(x.Name)
	x.Name = "Alex"
	fmt.Println(&x.Name)
	fmt.Println(x.Name)

}
