package dao

// Slot -
type Slot interface {
	SetNo(no int)
	GetNo() int
	SetVehicle(v Vehicle)
	GetVehicle() Vehicle
}
