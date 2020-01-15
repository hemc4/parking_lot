package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

func executeFile(path string) error{
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := splitCommand(scanner.Text())
		output, err := runCommand(command)
		if err != nil {
			fmt.Println(err.Error())
		}else{
			fmt.Println(output)
		}
	}
	return nil
}


