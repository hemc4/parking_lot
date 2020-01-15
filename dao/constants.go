package dao

import(
	"errors"
)

var (
	// ErrMaxSlotReached - 
	ErrMaxSlotReached = errors.New("Sorry, parking lot is full")
	// ErrCarNotFound - 
	ErrCarNotFound = errors.New("Not found")
)