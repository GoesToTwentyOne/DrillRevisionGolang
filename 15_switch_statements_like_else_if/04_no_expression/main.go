package main

import "fmt"

func main() {
	myGirlFriend := "candy"
	switch {
	case len(myGirlFriend) == 2:
		fmt.Println("myGirlFriend name onely 2 character")
	case myGirlFriend == "sandy":
		fmt.Println("My name is Sandy I'm 20 years old")
	case myGirlFriend == "mandy", myGirlFriend == "Alex":
		fmt.Println("That are my two girlfriend")
	case myGirlFriend == "Lname":
		fmt.Println("My name is Lname")
	case len(myGirlFriend) == 4, myGirlFriend == "candy":
		fmt.Println("This is Candy My Girlfriend and his name only 4 character")

	default:
		fmt.Println("Don't match the above")
	}
}
