package main 

import (
	"flag"
)

func main(){
	flag.Parse()
	// TODO create input handler inteface, like cmd, file, api, etc 
	if len(flag.Args()) > 0 {
		executeFile(flag.Args()[0])
	}else{
		executeInlineCommands()
	}
}