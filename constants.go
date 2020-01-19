package main

import (
	"errors"
)

var (
	// CmdPark -
	CmdPark = "park"
	// CmdCreateParkingLot -
	CmdCreateParkingLot = "create_parking_lot"
	// CmdStatus -
	CmdStatus = "status"
	// CmdLeave -
	CmdLeave = "leave"
	// CmdRegistrationNumberByColour -
	CmdRegistrationNumberByColour = "registration_numbers_for_cars_with_colour"
	// CmdSlotnoByCarColour -
	CmdSlotnoByCarColour = "slot_numbers_for_cars_with_colour"
	// CmdSlotnoByRegNumber -
	CmdSlotnoByRegNumber = "slot_number_for_registration_number"

	// ErrNoParkingLotCreated -
	ErrNoParkingLotCreated = errors.New("No parking lot avaible. Create parking lot first")
)
