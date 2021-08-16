package telegram

import (
	"fmt"
	"notify_bot/db"
	"strings"
)

func ParseCommands(data *messageFromTelegram) {
	var command string
	pos := 	strings.Index(data.Message.Text, "@")
	if pos == -1{
		command = data.Message.Text
	} else {
		if data.Message.Text[pos+1:] == BotName {
			command = data.Message.Text[:pos]
		}
	}
	switch command{
		case "/start": start(data.Message.Chat.ID)
	case "/help": help(data.Message.Chat.ID)
	case "/notify": notify(data.Message.Chat.ID)
	}
}

func IsCommand(data *messageFromTelegram) bool{
	if (data.Message.Entities[0]).Type == "bot_command"{
		return true
	}else {
		return false
	}
}

func start(ChatID int64){
	if !db.IsChatIDInDB(ChatID){
		token := GenerateNewToken()
		db.AddNewChat(ChatID, token)
	}
	message := "Send messages directly into the chat via http post. \n\n /notify — get webhook link for this chat."
	SendMessage(ChatID, message, "HTML", "")
}

func help(chatID int64){
	message := "Send notifications to chat easily One step integration. \n\n /notify_start — show webhook link for this chat"
	SendMessage(chatID, message, "HTML", "")
}

func notify(ChatID int64){
	if db.IsChatIDInDB(ChatID){
		token := db.GetToken(ChatID)
		message := fmt.Sprintf("Use this webhook for sending notifications to the chat:\n\n%s/%s\n\nMake a POST request with text in «message» param.", Url, token)
		SendMessage(ChatID, message, "HTML", "")
	}else {
		SendMessage(ChatID, "Print /start", "HTML", "")
	}
}
