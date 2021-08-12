package main

import (
	"fmt"
	"net/http"
	"notify_bot/telegram"
	"net/url"
)

var chatID int64

func handler(res http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    text := req.FormValue("message")
    parse_mode := req.FormValue("parse_mode")
    disable_web_page_preview := req.FormValue("disable_web_page_preview")
    if len(text) != 0{
    	telegram.SendMessage(chatID, text, parse_mode, disable_web_page_preview)
    } else {
    	dataFromChat := telegram.DecodeMessageFromJSON(req.Body)
    	chatID = dataFromChat.Message.Chat.ID
    	fmt.Println(dataFromChat.Message.Text)
    }
}

func makeHandle() string{
	u, err := url.Parse(telegram.Url)
	if err != nil{
	}
	handle := u.EscapedPath()
	return handle
}

func StartServer(port string){
	handle := makeHandle()
	http.HandleFunc(handle, handler)
	http.ListenAndServe("localhost:"+port, nil)
}
