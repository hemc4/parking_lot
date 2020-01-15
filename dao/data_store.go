package dao


// DataStore - 
type DataStore interface {
	Park(*Car) (*Slot, error)
	Leave(*Slot) (error)
	
	// Access
	GetAll() ([]*Slot, error)
	GetAllSlotsByColour(string) ([]*Slot, error)

	GetAllCarsByColour(string) ([]*Car, error)
	GetCarByRegNo(string) (*Car, error)
}