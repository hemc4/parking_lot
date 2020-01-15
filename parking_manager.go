package main

import(
	"fmt"
	"strings"
	"strconv"
	dao "parking_lot/dao"
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
	// CmdRegistrationNumberByColor - 
	CmdRegistrationNumberByColor = "registration_numbers_for_cars_with_colour"
	// CmdSlotnoByCarColor - 
	CmdSlotnoByCarColor = "slot_numbers_for_cars_with_colour"
	// CmdSlotnoByRegNumber - 
	CmdSlotnoByRegNumber = "slot_number_for_registration_number"
)

var db *dao.InMemoryDB

func createParkingLot(maxSlots int) (string, error){
	db = dao.NewInMemoryDB(maxSlots)
	out := fmt.Sprintf("Created a parking lot with %d slots", maxSlots)
	return out, nil
}

func park(regNo, colour string) (string, error) {
	c := dao.NewCar(regNo, colour)
	slot, err := db.Park(c)
	var out string
	if err != nil {
		out = err.Error()
	}else{
		out = fmt.Sprintf("Allocated slot number: %d", slot.No)
	}
	return out, err
}

func status() (string, error) {
	slotsList, err := db.GetAll()
	var out string
	if err != nil {
		out = err.Error()
	}else{
		out = fmt.Sprintf("Slot No.\tRegistration No\tColor")
		for _, slot := range slotsList {
			if slot!= nil && slot.Car != nil {
				out += fmt.Sprintf("\n%d %s %s", slot.No, slot.Car.RegNo, slot.Car.Colour)
			}
		}
	}
	return out, err
}

func runCommand(command []string) (string, error) {
	switch command[0] {
	case CmdCreateParkingLot:
		maxSlots, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err.Error())
		}
		return createParkingLot(maxSlots)
	case CmdPark:
		return park(command[1], command[2])
	case CmdStatus:
		return status()
	default:
	}
	return "", nil
}


// TODO create a struct for command with args and parse all strings to that command,
// just like commmand pattern
func splitCommand(command string) []string {
	splitCommand := []string{}

	for _, s := range strings.Split(command, " ") {
		if s != "" {
			splitCommand = append(splitCommand, s)
		}
	}
	return splitCommand
}