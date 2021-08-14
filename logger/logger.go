package logger

import (
	log "github.com/sirupsen/logrus"
)

func Info(info string){
	log.Info(info)

}

func Warn(text string){
	log.Warn(text)
}

func Fatal(text string){
	log.Fatal(text)
}


