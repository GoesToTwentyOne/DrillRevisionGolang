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

func main() {
	hatchback := car{}
	sedan := car{}
	suv := car{}
	crossover := car{}
	coupe := car{}
	convertible := car{}
	cars := []car{hatchback, sedan, suv, crossover, coupe, convertible}
	fmt.Println(cars)
	fmt.Println("-------------------------------------------------")
	fishing := boat{}
	house := boat{}
	bowrider := boat{}
	game := boat{}
	lifeboat := boat{}
	catamaran := boat{}
	boats := []boat{fishing, house, bowrider, game, lifeboat, catamaran}
	fmt.Println(boats)
	fmt.Println("-------------------------------------------------")

	piston := plane{}
	airbus := plane{}
	concorde := plane{}
	military := plane{}
	fighter := plane{}
	transport := plane{}
	planes := []plane{piston, airbus, concorde, military, fighter, transport}
	fmt.Println(planes)
	fmt.Println("-------------------------------------------------")

	for key, values := range cars {
		fmt.Println(key, "  -  ", values)
	}
	fmt.Println("-------------------------------------------------")
	for key, values := range boats {
		fmt.Println(key, "  -  ", values)
	}
	fmt.Println("-------------------------------------------------")
	for key, values := range planes {
		fmt.Println(key, "  -  ", values)
	}
	fmt.Println("-------------------------------------------------")

}
