package structs

import (
	"encoding/json"
	"io"
	)
type messageFromTelegram struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type messageToTelegram struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type dataForWebhook struct{
	Url string `json:"url"`
}

func DecodeMessageFromJSON(data io.Reader) *messageFromTelegram{
	message := &messageFromTelegram{}
	if err := json.NewDecoder(data).Decode(message); err != nil {
	}
	return message
}

func СodeMessageToJSON(chatID int64, text string, Parse_mode string) []byte{
	message := &messageToTelegram{ ChatID: chatID, Text: text, ParseMode: Parse_mode, }
	messageBytes, err := json.Marshal(message)
	if err != nil {
	}
	return messageBytes
}

func MakeDataForWebhook(url string) []byte{
	data := &dataForWebhook{
		Url: url,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
	}
	return dataBytes
}
