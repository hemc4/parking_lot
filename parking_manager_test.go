package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	dao "parking_lot/dao"
)


func TestCreateParkingLot(t *testing.T) {
	var cases = []struct {
		input int
		output string
		err error
	}{
		{0, "", dao.ErrInvalidMaxSlots},
		{6, "Created a parking lot with 6 slots", nil},
	}

	for _, testCase := range cases {
		recievedOutput, err := createParkingLot(testCase.input)
		if testCase.err != nil {
			assert.Equal(t, testCase.err, err, "they should be equal")
		} else {
			assert.Equal(t, testCase.output, recievedOutput, "they should be equal")
		}
	}
}


func TestSplitCommand(t *testing.T) {
	var cases = []struct {
		input string
		output []string
	}{
		{"create_parking_lot 6", []string{"create_parking_lot", "6"}},
	}

	for _, testCase := range cases {
		recievedOutput := splitCommand(testCase.input)
		assert.Equal(t, testCase.output, recievedOutput, "they should be equal")
		
	}
}


func TestRunCommand(t *testing.T) {
	var cases = []struct {
		input []string
		output string
		err error
	}{
		{[]string{"create_parking_lot", "6"}, "Created a parking lot with 6 slots", nil},
		{[]string{"create_multi_story_parking_lot", "3", "6"}, "", ErrCommandNotSupported},
	}

	for _, testCase := range cases {
		recievedOutput, err := runCommand(testCase.input)
		if testCase.err != nil {
			assert.Equal(t, testCase.err, err, "they should be equal")
		} else {
			assert.Equal(t, testCase.output, recievedOutput, "they should be equal")
		}
	}
}