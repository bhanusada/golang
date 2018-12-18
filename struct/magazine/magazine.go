package magazine

//Subscriber type to set subscriber details
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

//Employee type to set employee details
type Employee struct {
	Name   string
	Salary float64
	Address
}

//Address type to setaddress details
type Address struct {
	Flatno  string
	City    string
	State   string
	Country string
}
