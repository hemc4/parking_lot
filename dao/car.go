package dao

// Car - 
type Car struct {
	RegNo  string 
	Colour string 
}

// NewCar - constuctor
func NewCar(regNo, colour string) *Car {
	return &Car{
		RegNo: regNo,
		Colour: colour,
	}
}
