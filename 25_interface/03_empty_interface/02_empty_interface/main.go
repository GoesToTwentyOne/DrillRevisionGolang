package main

import "fmt"

type vehicle struct {
	Seat     int
	MaxSpeed int
	color    string
}
type car struct {
	vehicle
	wheel int
	dors  int
}
type boat struct {
	vehicle
	Lenght    int
	Container int
}
type plane struct {
	vehicle
	Jtair int
	Jet   bool
}
type vehicles interface {
}

func main() {
	hatchback := car{}
	sedan := car{}
	suv := car{}
	crossover := car{}
	coupe := car{}
	convertible := car{}
	fishing := boat{}
	house := boat{}
	bowrider := boat{}
	game := boat{}
	lifeboat := boat{}
	catamaran := boat{}
	piston := plane{}
	airbus := plane{}
	concorde := plane{}
	military := plane{}
	fighter := plane{}
	transport := plane{}
	rides := []vehicles{hatchback, sedan, suv, crossover, coupe, convertible, fishing, house, bowrider, game,
		lifeboat, catamaran, piston, airbus, concorde, military, fighter, transport}

	fmt.Println(rides)
	for key, values := range rides {
		fmt.Println(key, "   -   ", values)
	}

}
