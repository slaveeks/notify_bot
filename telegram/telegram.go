package telegram

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"notify_bot/logger"
	"time"
)

var TelegramUrl = "https://api.telegram.org/bot"
var Token string
var Url string
var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func SendMessage(chatID int64, text string, parseMode string, disableWebPagePreview string){
	if parseMode == ""{
    		parseMode = "None"
        }
	if disableWebPagePreview == ""{
    		disableWebPagePreview = "false"
    	}
	reqMessage := СodeMessageToJSON(chatID, text, parseMode, disableWebPagePreview)
	_, err := http.Post(fmt.Sprintf("%s%s/sendMessage", TelegramUrl, Token), "application/json", bytes.NewBuffer(reqMessage))
	if err != nil {
		logger.Warn("Error while sending message")
	}
	logger.Info("Message is send to chat")
}

func SetWebhook(){
	reqData := MakeDataForWebhook(Url)
	_, err := http.Post(fmt.Sprintf("%s%s/setWebhook", TelegramUrl, Token), "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		logger.Fatal("Webhook does not be set")
	}
	logger.Info(fmt.Sprintf("Webhook was set to ", Url))
}

func GenerateNewToken() string {
	rand.Seed(time.Now().UnixNano())
	token := make([]rune, 8)
	for i := range token {
		token[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(token)
}
