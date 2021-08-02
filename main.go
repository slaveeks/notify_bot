package main

import (
	"os"
	"log"
	"flag"
	"fmt"
)

func main() {
	mySet := flag.NewFlagSet("start_bot",flag.ExitOnError)

	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: <command> <arguments>\nCommand: start_bot\nArguments:\n Url prefix, Listen Port, telegram token\n")
		os.Exit(0)
	}
	err := mySet.Parse(os.Args)
	if err != nil {
		log.Fatalf("Invalid number of arguments")
	}
	
}
