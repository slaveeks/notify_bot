package main

import (
	"flag"
	"fmt"
	"notify_bot/db"
	"notify_bot/telegram"
	"os"
)

func main() {
	token := flag.String("token", "", "Telegram bot token")
	addr := flag.String("addr", "localhost:1330", "Server's address")
	mongoUrl := flag.String("mongo", "mongodb://127.0.0.1:27017", "Url for Database")
	url := flag.String("url", "", "Url for telegram callbacks")
	cert := flag.String("cert", "", "Path for certificate")
	key := flag.String("key", "", "Path for key")
	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: %s\n\nArguments:\n ", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	telegram.Token = *token
	telegram.Url = *url
	db.InitDB(*mongoUrl)
	if len(*cert) != 0 && len(*key) !=0 {
		telegram.SetWebhook(*cert)
		StartServerWithTLS(*addr, *key, *cert)
	} else {
		telegram.SetWebhook("")
		StartServer(*addr)
	}
	StartServer(*addr)
}
