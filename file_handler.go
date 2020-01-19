package main

import (
	"bufio"
	"fmt"
	"os"
)

func executeFile(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := splitCommand(scanner.Text())
		//fmt.Println(command)
		output, err := runCommand(command)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(output)
		}
	}
	return nil
}
