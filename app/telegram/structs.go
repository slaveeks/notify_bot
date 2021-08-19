package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"notify_bot/logger"
)

type User struct{
	Result struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		UserName  string `json:"username"`
	} `json:"result"`
}

type MessageEntity struct {
	Type string `json:"type"`
	Offset int `json:"offset"`
	Length int `json:"length"`
}

type messageFromTelegram struct {
	Message struct {
		Entities []MessageEntity `json:"entities"`
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type messageToTelegram struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview string `json:"disable_web_page_preview"`
}

type dataForWebhook struct{
	Url string `json:"url"`
	cert string `json:"certificate"`
}

func DecodeMessageFromJSON(data io.Reader) *messageFromTelegram{
	message := &messageFromTelegram{}
	if err := json.NewDecoder(data).Decode(message); err != nil {
		logger.Warn("Error while decoding message")
	}
	return message
}

func СodeMessageToJSON(chatID int64, text string, ParseMode string, disableWebPagePreview string) []byte{
	message := &messageToTelegram{ ChatID: chatID, Text: text, ParseMode: ParseMode, DisableWebPagePreview: disableWebPagePreview, }
	messageBytes, err := json.Marshal(message)
	if err != nil {
		logger.Warn(fmt.Sprint("Error while coding message\n", err))
	}
	return messageBytes
}

func MakeDataForWebhook(url string, cert string) []byte{
	data := &dataForWebhook{}
	if len(cert) != 0 {
		data = &dataForWebhook{
			Url:  url,
			cert: "@" + cert,
		}
	}else {
		data = &dataForWebhook{
			Url: url,
		}
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		logger.Warn(fmt.Sprint("Error while coding message\n", err))
	}
	return dataBytes
}

func ConvertToUser(data io.Reader) *User {
	var Bot *User
	body, err := ioutil.ReadAll(data)
	if err != nil {
		logger.Warn(fmt.Sprint(err))
	}
	err = json.Unmarshal(body, &Bot)
	if err != nil {
		logger.Warn(fmt.Sprint(err))
	}
	return Bot
}


