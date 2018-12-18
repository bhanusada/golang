package main

import (
	"fmt"

	"github.com/bhanu/struct/magazine"
)

/*type structType struct {
	name   string
	rate   float64
	active bool
}*/

func main() {
	mystruct, stringPointer := defaultValues()
	fmt.Println(*mystruct)
	fmt.Println(*stringPointer)
	changeRate(mystruct)
	fmt.Println(*mystruct)
}

func defaultValues() (*magazine.Subscriber, *string) {
	//var s magazine.Subscriber
	var s1 string
	s1 = "hjgfhgfh"
	//s.Name = "bhanu"
	//s.Rate = 345.45
	//s.Active = false
	s := magazine.Subscriber{
		Name:   "bhanu",
		Rate:   23.34,
		Active: false,
		Address: magazine.Address{
			Flatno:  "1014",
			City:    "Smyrna",
			State:   "GA",
			Country: "USA"}}
	return &s, &s1
}

func changeRate(s *magazine.Subscriber) {
	s.Rate = 23.4
}
