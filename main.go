package main 

import (
	"flag"
)

func main(){
	flag.Parse()
	// TODO Convert it to multi(cmd, file, api ...) support input/output handler inteface 
	if len(flag.Args()) > 0 {
		executeFile(flag.Args()[0])
	}else{
		executeInlineCommands()
	}
}