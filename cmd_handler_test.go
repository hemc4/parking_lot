package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
)

func TestRunCmdInput(t *testing.T) {
	var cases = []struct {
		input  string
		output string
	}{
		{"create_parking_lot 6", "Created a parking lot with 6 slots\n"},
		{"create_parking_lot  6", "Created a parking lot with 6 slots\n"},
		{"create_parking_lot    6", "Created a parking lot with 6 slots\n"},
	}

	for _, testCase := range cases {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		runCmdInput(testCase.input)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		assert.Equal(t, testCase.output, string(out), "they should be equal")
	}
}