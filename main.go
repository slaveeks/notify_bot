package main

import (
	"flag"
	"fmt"
	"log"
	"notify_bot/telegram"
	"os"
)

func main() {
	mySet := flag.NewFlagSet("start_bot",flag.ExitOnError)

	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: <command> <arguments>\nCommand: start_bot\nArguments:\n Url prefix, Listen Port, telegram token")
		os.Exit(0)
	}
	err := mySet.Parse(os.Args)
	if err != nil {
		log.Fatalf("Invalid number of arguments")
	}
	url := os.Args[2]
	port := os.Args[3]
	token := os.Args[4]
	telegram.Token = token
	telegram.Url = url
	telegram.SetWebhook()
	StartServer(port)
}
