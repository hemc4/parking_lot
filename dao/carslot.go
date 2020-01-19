package dao

// CarSlot -
type CarSlot struct {
	no      int
	vehicle Vehicle
}

// NewSlot - constuctor
func NewSlot(no int) *CarSlot {
	s := new(CarSlot)
	s.no = no
	return s
}

//SetNo -
func (s *CarSlot) SetNo(no int) {
	s.no = no
}

//GetNo -
func (s *CarSlot) GetNo() int {
	return s.no
}

//SetVehicle -
func (s *CarSlot) SetVehicle(v Vehicle) {
	s.vehicle = v
}

//GetVehicle -
func (s *CarSlot) GetVehicle() Vehicle {
	return s.vehicle
}
