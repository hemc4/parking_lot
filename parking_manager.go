package main

import(
	"fmt"
	"strings"
	"strconv"
	dao "parking_lot/dao"
)

var db dao.DataStore

func createParkingLot(maxSlots int) (string, error){
	db = dao.NewInMemoryDB(maxSlots)
	out := fmt.Sprintf("Created a parking lot with %d slots", maxSlots)
	return out, nil
}

func park(regNo, colour string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
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

func leave(no int) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	s := dao.NewSlot(no)
	err := db.Leave(s)
	var out string
	if err != nil {
		out = err.Error()
	}else{
		out = fmt.Sprintf("Slot number %d is free", no)
	}
	return out, err
}


func status() (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slotsList, err := db.GetAll()
	var out string
	if err != nil {
		out = err.Error()
	}else{
		out = fmt.Sprintf("Slot No.    Registration No    Colour")
		for _, slot := range slotsList {
			if slot!= nil && slot.Car != nil {
				out += fmt.Sprintf("\n%d           %s      %s", slot.No, slot.Car.RegNo, slot.Car.Colour)
			}
		}
	}
	return out, err
}

func getAllRegNoBycolour(colour string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slotsList, err := db.GetAllSlotsByColour(colour)
	var out string
	if err != nil {
		out = err.Error()
	}else{
		first := true
		for _, slot := range slotsList {
			if slot!= nil && slot.Car != nil {
				if first {
					out += fmt.Sprintf("%s",slot.Car.RegNo)
				}else{
					out += fmt.Sprintf(", %s",slot.Car.RegNo)
				}
				first = false;
			}
		}
	}
	return out, err

}

func getAllSlotNoBycolour(colour string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slotsList, err := db.GetAllSlotsByColour(colour)
	var out string
	if err != nil {
		out = err.Error()
	}else{
		first := true
		for _, slot := range slotsList {
			if slot!= nil {
				if first {
					out += fmt.Sprintf("%d", slot.No)
				}else{
					out += fmt.Sprintf(", %d", slot.No)
				}
				first = false;
			}
		}
	}
	return out, err
}

func getSlotNoByRegNo(regNo string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slot,err := db.GetSlotByRegNo(regNo)
	var out string
	if err != nil {
		out = err.Error()
	}else{
		out = fmt.Sprintf("%d", slot.No)
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
	case CmdLeave:
		no, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err.Error())
		}
		return leave(no)
	case CmdStatus:
		return status()
	case CmdRegistrationNumberByColour:
		return getAllRegNoBycolour(command[1])
	case CmdSlotnoByCarColour:
		return getAllSlotNoBycolour(command[1])
	case CmdSlotnoByRegNumber:
		return getSlotNoByRegNo(command[1])
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