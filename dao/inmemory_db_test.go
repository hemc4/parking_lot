package dao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInMemoryDB(t *testing.T) {
	var cases = []struct {
		input  int
		output *InMemoryDB
		err    error
	}{
		{0, nil, ErrInvalidMaxSlots},
		{6, new(InMemoryDB), nil},
	}

	for _, testCase := range cases {
		recievedOutput, err := NewInMemoryDB(testCase.input)
		if testCase.err != nil {
			assert.Equal(t, testCase.err, err, "they should be equal")
		} else {
			if assert.NotNil(t, recievedOutput.slots) {
				assert.Equal(t, testCase.input, len(recievedOutput.slots), "they should be equal")
			}
		}
	}
}

func TestGetNextEmptySlot(t *testing.T) {
	db, _ := NewInMemoryDB(4)
	slot, _ := db.GetNextEmptySlot()
	if assert.NotNil(t, slot) {
		assert.Equal(t, 1, slot.No, "they should be equal")
	}
	car := NewCar("KA-01", "White")
	db.Park(car)

	slot2, _ := db.GetNextEmptySlot()
	if assert.NotNil(t, slot2) {
		assert.Equal(t, 2, slot2.No, "they should be equal")
	}

	db.Leave(slot)

	nexSlot, _ := db.GetNextEmptySlot()
	if assert.NotNil(t, nexSlot) {
		assert.Equal(t, 1, nexSlot.No, "they should be equal")
	}

}
