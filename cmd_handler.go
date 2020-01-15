package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func executeInlineCommands() error{
	exitCommand := false
	var buffReader *bufio.Reader
	buffReader = bufio.NewReader(os.Stdin)

	for !exitCommand {
		inputText, _ := buffReader.ReadString('\n')
		inputText = strings.TrimRight(inputText, "\r\n")
		if inputText == "exit" {
			break
		}
		command := splitCommand(inputText)
		//fmt.Println(command)
		output, err := runCommand(command)
		if err != nil {
			fmt.Println(err.Error())
		}else{
			fmt.Println(output)
		}
	}
	return nil
}
