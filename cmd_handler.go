package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func executeInlineCommands() error {
	exitCommand := false
	var buffReader *bufio.Reader
	buffReader = bufio.NewReader(os.Stdin)

	for !exitCommand {
		inputText, _ := buffReader.ReadString('\n')
		inputText = strings.TrimRight(inputText, "\r\n")
		if inputText == "exit" {
			break
		}
		runCmdInput(inputText)
	}
	return nil
}

func runCmdInput(inputText string) {
	command := splitCommand(inputText)
	//fmt.Println(command)
	output, err := runCommand(command)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(output)
	}
}
