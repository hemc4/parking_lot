package dao

// Slot -
type Slot struct {
	no      int
	vehicle Vehicle
}

// NewSlot - constuctor
func NewSlot(no int) *Slot {
	s := new(Slot)
	s.no = no
	return s
}

//SetNo -
func (s *Slot) SetNo(no int) {
	s.no = no
}

//GetNo -
func (s *Slot) GetNo() int {
	return s.no
}

//SetVehicle -
func (s *Slot) SetVehicle(v Vehicle) {
	s.vehicle = v
}

//GetVehicle -
func (s *Slot) GetVehicle() Vehicle {
	return s.vehicle
}
