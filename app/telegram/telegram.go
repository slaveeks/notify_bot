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
var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var BotName string

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
		logger.Warn(fmt.Sprint("Error while sending message", err))
	}
	logger.Info("Message is send to chat")
}

func SetWebhook(cert string){
	reqData := MakeDataForWebhook(Url, cert)
	_, err := http.Post(fmt.Sprintf("%s%s/setWebhook", TelegramUrl, Token), "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		logger.Fatal(fmt.Sprint("Webhook does not be set\n", err))
	}
	logger.Info(fmt.Sprintf("Webhook was set to ", Url))
	BotName = GetBotName()
}


func GetBotName() string{
	resp, err := http.Get(fmt.Sprintf("%s%s/getMe", TelegramUrl, Token))
	if err != nil {
		logger.Fatal(fmt.Sprint("Error while getting information about bot\n", err))
	}
	bot := ConvertToUser(resp.Body)
	logger.Info(fmt.Sprint(bot))
	return bot.Result.UserName
}

func GenerateNewToken() string {
	rand.Seed(time.Now().UnixNano())
	token := make([]rune, 8)
	for i := range token {
		token[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(token)
}


