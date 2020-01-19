package dao

import(
	"errors"
)

var (
	// ErrMaxSlotReached - 
	ErrMaxSlotReached = errors.New("Sorry, parking lot is full")
	// ErrCarNotFound - 
	ErrCarNotFound = errors.New("Not found")
	// ErrInvalidMaxSlots - 
	ErrInvalidMaxSlots = errors.New("Max slots should be greter than 0")
)