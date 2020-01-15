package dao

// Slot - 
type Slot struct {
	No  	int
	Car		*Car
}

// NewSlot - constuctor
func NewSlot(no int) *Slot {
	s := new(Slot)
	s.No = no
	return s
}