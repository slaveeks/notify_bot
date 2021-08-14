package telegram

import (
	"fmt"
	"notify_bot/db"
)

func ParseCommands(data *messageFromTelegram) {
	switch data.Message.Text{
		case "/start": start(data.Message.Chat.ID)
	case "/help": help()
	case "/notify": notify(data.Message.Chat.ID)
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

func help(){

}

func notify(ChatID int64){
	if db.IsChatIDInDB(ChatID){
		token := db.GetToken(ChatID)
		fmt.Println(token)
		message := fmt.Sprintf("Use this webhook for sending notifications to the chat:\n\n%s/%s\n\nMake a POST request with text in «message» param.", Url, token)
		SendMessage(ChatID, message, "HTML", "")
	}else {
		SendMessage(ChatID, "Print /start", "HTML", "")
	}
}
