package main

import "fmt"

func main() {
	love := 304
	fmt.Println(&love) //304
	changeMe(&love)    //pass address
	fmt.Println(&love) //address
	fmt.Println(love)  //value 618

}
func changeMe(n *int) {
	fmt.Println(n)
	fmt.Println(*n)
	*n = 618
	fmt.Println(n)
	fmt.Println(*n)

}
