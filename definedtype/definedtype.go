package main

import (
	"fmt"
)

// Gallons type
type Gallons float64

// Liters type
type Liters float64

//MyType type
type MyType string

//MyNumber type
type MyNumber struct {
	Number int
}

func main() {
	var fuel Gallons
	fuel = 3.24
	liters := Liters(fuel * 3.7)
	fmt.Printf("Fuel in liters %.2f\n", liters)

	value := MyType("MyType value")
	value.sayHi()

	number := MyNumber{Number: 2}
	number.double()
	number.display()
}

func (m MyType) sayHi() {
	fmt.Printf("Hi from %s\n", m)
}

func (n *MyNumber) double() {
	//fmt.Println(n.Number)
	n.Number *= 2
}

func (n MyNumber) display() {
	fmt.Printf("The number is %d\n", n.Number)
}
