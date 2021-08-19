package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"notify_bot/db"
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
	data := c.Request().Body
	logger.Infof("%s", data)
	dataFromChat := telegram.DecodeMessageFromJSON(data)
	logger.Info(fmt.Sprint(dataFromChat))
	if telegram.IsCommand(dataFromChat) {
		telegram.ParseCommands(dataFromChat)
	}
	return nil
}

func makeHandle() string{
	u, err := url.Parse(telegram.Url)
	if err != nil{
		logger.Fatal(fmt.Sprint("Url was not parsed\n", err))
	}
	handle := u.EscapedPath()
	logger.Info(fmt.Sprintf("Route is set to /%s", handle))
	return handle
}

func StartServer(addr string){
	handle := makeHandle()
	e := echo.New()
	e.POST(handle, telegramCallbacks)
	e.POST(handle+"/:token", WebhookHandler)
	e.Logger.Fatal(e.Start(addr))
}

func StartServerWithTLS(addr string, key string, cert string){
	handle := makeHandle()
	e := echo.New()
	e.POST(handle, telegramCallbacks)
	e.POST(handle+"/:token", WebhookHandler)
	if err := e.StartTLS(":8443", cert, key); err != http.ErrServerClosed {
		logger.Fatal(err)
	}
}
