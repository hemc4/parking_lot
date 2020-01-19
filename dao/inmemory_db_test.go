package dao

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestNewInMemoryDB(t *testing.T) {
	var cases = []struct {
		input int
		output *InMemoryDB
		err error
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