package dao

import(
	"errors"
)

var (
	// ErrMaxSlotReached - 
	ErrMaxSlotReached = errors.New("Sorry, parking lot is full")
	// ErrNoCarsParked - 
	ErrNoCarsParked = errors.New("No cars parked")
	// ErrCarNotFound - 
	ErrCarNotFound = errors.New("Not found")
	// ErrCarWithColourNotFound - 
	ErrCarWithColourNotFound = errors.New("Car with specified colour not found")
)