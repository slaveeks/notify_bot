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
    if len(text) != 0{
    	telegram.SendMessage(chatID, text, parse_mode)
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
