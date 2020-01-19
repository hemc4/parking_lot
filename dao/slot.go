package dao

// Slot -
type Slot struct {
	No      int
	Vehicle Vehicle
}

// NewSlot - constuctor
func NewSlot(no int) *Slot {
	s := new(Slot)
	s.No = no
	return s
}
