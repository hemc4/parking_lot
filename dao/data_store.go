package dao

// DataStore -
type DataStore interface {
	//modifiers
	Park(*Car) (*Slot, error)
	Leave(*Slot) error
	//readers
	GetAll() ([]*Slot, error)
	GetAllSlotsByColour(string) ([]*Slot, error)
	GetSlotByRegNo(string) (*Slot, error)
}
