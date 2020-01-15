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
	{"create_parking_lot 6", "Created a parking lot with 6 slots\n"},
	{"park KA­-01-­HH­-1234 White", "Allocated slot number: 1\n"},
	{"park KA-­01-­HH-­9999 White", "Allocated slot number: 2\n"},
	{"park KA­-01-­BB-­0001 Black", "Allocated slot number: 3\n"},
	{"park KA-­01­-HH­-7777 Red", "Allocated slot number: 4\n"},
	{"park KA-­01­-HH­-2701 Blue", "Allocated slot number: 5\n"},
	{"park KA-­01­-HH­-3141 Black", "Allocated slot number: 6\n"},
	{"leave 4", "Slot number 4 is free\n"},
	{"status", `Slot No.	Registration No	Colour
1		KA­-01-­HH­-1234	White
2		KA-­01-­HH-­9999	White
3		KA­-01-­BB-­0001	Black
5		KA-­01­-HH­-2701	Blue
6		KA-­01­-HH­-3141	Black` + "\n"},
	{"park KA-­01-­P­-333 White", "Allocated slot number: 4\n"},
	{"park DL-­12­-AA-­9999 White", "Sorry, parking lot is full\n"},
	{"registration_numbers_for_cars_with_colour White", "KA­-01-­HH­-1234, KA-­01-­HH-­9999, KA-­01-­P­-333\n"},
	{"slot_numbers_for_cars_with_colour White", "1, 2, 4\n"},
	{"slot_number_for_registration_number KA-­01­-HH­-3141", "6\n"},
	{"slot_number_for_registration_number MH­-04-­AY-­1111", "Not found\n"},
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