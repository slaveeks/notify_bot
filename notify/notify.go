package notify

import (
	"../structs"
	"../telegram"
	"strings"
	"net/http"
	"fmt"
)

var chatID int64

func handler(res http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    text := req.FormValue("message")
    parse_mode := req.FormValue("parse_mode")
    disable_web_page_preview := req.FormValue("disable_web_page_preview")
    if parse_mode == ""{
    	parse_mode = "None"
    }
    if disable_web_page_preview == ""{
    disable_web_page_preview = "false"
    }
    if len(text) != 0{
    	telegram.SendMessage(chatID, text, parse_mode, disable_web_page_preview)
    } else {
    	dataFromChat := structs.DecodeMessageFromJSON(req.Body)
    	chatID = dataFromChat.Message.Chat.ID
    	fmt.Println(dataFromChat.Message.Text)
    }
}


func makeHandle() string{
	urlArray := strings.Split(telegram.Url, "/")
	handle := "/"
	for i:=3 ; i<len(urlArray); i++{
		handle += urlArray[i] + "/"
	}
	return handle
}

func StartServer(port string){
	handle := makeHandle()
	http.HandleFunc(handle, handler)
	http.ListenAndServe("localhost:"+port, nil)
}
