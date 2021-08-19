package main

import (
	"flag"
	"fmt"
	"notify_bot/db"
	"notify_bot/logger"
	"notify_bot/telegram"
	"os"
)

func main() {
	var url string
	var mongoUrl string
	var addr string
	var token string
	mySet := flag.NewFlagSet("",flag.ExitOnError)
	mySet.StringVar(&url, "url", "", "Url for webhooks")
	mySet.StringVar(&mongoUrl, "mongo", "mongodb://127.0.0.1:27017", "Url for Database")
	mySet.StringVar(&addr, "addr", "localhost:1330", "Server's address")
	mySet.StringVar(&token, "token", "", "Telegram bot token")
	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n\nArguments:\n ", os.Args[0])
		mySet.PrintDefaults()
		os.Exit(0)
	}
	err := mySet.Parse(os.Args[1:])
	if err != nil {
		logger.Fatal("Invalid number of arguments")
	}
	telegram.Token = token
	telegram.Url = url
	db.InitDB(mongoUrl)
	telegram.SetWebhook()
	StartServer(addr)
}
