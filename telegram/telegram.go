package telegram

import (
	"notify_bot/structs"
	"bytes"
	"net/http"
	"fmt"
)

var TelegramUrl = "https://api.telegram.org/bot"
var Token string
var Url string

func SendMessage(chatID int64, text string, parse_mode string, disable_web_page_preview string){
	if parse_mode == ""{
    		parse_mode = "None"
        }
	if disable_web_page_preview == ""{
    		disable_web_page_preview = "false"
    	}
	reqMessage := structs.СodeMessageToJSON(chatID, text, parse_mode, disable_web_page_preview)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TelegramUrl, Token), "application/json", bytes.NewBuffer(reqMessage))
	if err != nil {
	}
}

func SetWebhook(){
	reqData := structs.MakeDataForWebhook(Url)
	_, err := http.Post(fmt.Sprintf("%s%s/setWebhook", TelegramUrl, Token), "application/json", bytes.NewBuffer(reqData))
	if err != nil {
	}
}
