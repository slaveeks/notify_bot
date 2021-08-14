package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/url"
	"notify_bot/db"
	"notify_bot/logger"
	"notify_bot/telegram"
)

func WebhookHandler(c echo.Context) error{
	token := c.Param("token")
	if db.IsTokenInDB(token) {
		text := c.FormValue("message")
		parseMode := c.FormValue("parse_mode")
		disableWebPagePreview := c.FormValue("disable_web_page_preview")
		chatID := db.GetChatID(token)
		if len(text) != 0 {
			telegram.SendMessage(chatID, text, parseMode, disableWebPagePreview)
		}
	}else{
		logger.Info("Invalid Token")
	}
	return nil
}

func telegramCallbacks(c echo.Context) error {
	dataFromChat := telegram.DecodeMessageFromJSON(c.Request().Body)
	telegram.ParseCommands(dataFromChat)
	return nil
}

func makeHandle() string{
	u, err := url.Parse(telegram.Url)
	if err != nil{
		logger.Fatal("Url was not parsed")
	}
	handle := u.EscapedPath()
	logger.Info(fmt.Sprintf("Route is set to %s", handle))
	return handle
}

func StartServer(port string){
	handle := makeHandle()
	e := echo.New()
	e.POST(handle, telegramCallbacks)
	e.POST(handle+"/:token", WebhookHandler)
	e.Logger.Fatal(e.Start(":"+port))
	logger.Info(fmt.Sprintf("Sever was started on %s", port))
}
