package main

import (
	"io/ioutil"
	"os"
	"testing"
)

var parkingCommands = []struct {
	input  string
	output string
}{
	{"create_parking_lot 6", "Created a parking lot with 6 slots"},
	{"park KA­-01-­HH­-1234 White", "Allocated slot number: 1"},
	{"park KA-­01-­HH-­9999 White", "Allocated slot number: 2"},
	{"park KA­-01-­BB-­0001 Black", "Allocated slot number: 3"},
	{"park KA-­01­-HH­-7777 Red", "Allocated slot number: 4"},
	{"park KA-­01­-HH­-2701 Blue", "Allocated slot number: 5"},
	{"park KA-­01­-HH­-3141 Black", "Allocated slot number: 6"},
	{"leave 4", "Slot number 4 is free"},
	{"status", `Slot No.	Registration No	Colour
1		KA­-01-­HH­-1234	White
2		KA-­01-­HH-­9999	White
3		KA­-01-­BB-­0001	Black
5		KA-­01­-HH­-2701	Blue
6		KA-­01­-HH­-3141	Black`},
	{"park KA-­01-­P­-333 White", "Allocated slot number: 4"},
	{"park DL-­12­-AA-­9999 White", "Sorry, parking lot is full"},
	{"registration_numbers_for_cars_with_colour White", "KA­-01-­HH­-1234, KA-­01-­HH-­9999, KA-­01-­P­-333"},
	{"slot_numbers_for_cars_with_colour White", "1, 2, 4"},
	{"slot_number_for_registration_number KA-­01­-HH­-3141", "6"},
	{"slot_number_for_registration_number MH­-04-­AY-­1111", "Not found"},
}

func TestRunCommands(t *testing.T) {

	for _, parkingCommand := range parkingCommands {

		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		runCmdInput(parkingCommand.input)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		if parkingCommand.output != string(out) {
			t.Errorf("Invalid response %q =>  want \n %q got \n %q", parkingCommand.input, parkingCommand.output, out)
		}
	}
}