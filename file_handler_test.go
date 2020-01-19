package main

import (
	"errors"
	"testing"
)

func TestExecuteFile(t *testing.T) {
	var cases = []struct {
		input  string
		output error
	}{
		{"functional_spec/fixtures/file_input.txt", nil},
		{"dummy.txt", errors.New("Path Error")},
	}
	for _, testCase := range cases {
		err := executeFile(testCase.input)
		if testCase.output == nil && err != nil {
			t.Errorf("Invalid response %q =>  want %q got %q", testCase.input, testCase.output, err.Error())
		}
	}
}
