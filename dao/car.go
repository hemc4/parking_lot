package dao

// Car -
type Car struct {
	regNo  string
	colour string
}

// NewCar - constuctor
func NewCar(regNo, colour string) *Car {
	return &Car{
		regNo:  regNo,
		colour: colour,
	}
}

// GetRegNo -
func (c *Car) GetRegNo() string {
	return c.regNo
}

// GetColour -
func (c *Car) GetColour() string {
	return c.colour
}
