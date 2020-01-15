package main 

import (
	"fmt"
	"flag"
)

func main(){
	fmt.Println("hello ")
	// TODO create input handler inteface, like cmd, file, api, etc 
	if len(flag.Args()) > 0 {
		executeFile(flag.Args()[0])
	}else{
		executeInlineCommands()
	}
}