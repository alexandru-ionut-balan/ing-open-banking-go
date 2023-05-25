package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()
var sugar = logger.Sugar()

const zapSyncErrorMessage = "Zap logger cannot flush its buffer!"

func Info(args ...any) {
	sugar.Infoln(args)
	if err := sugar.Sync(); err != nil {
		log.Println(zapSyncErrorMessage)
	}
}

func Warn(args ...any) {
	sugar.Warnln(args)
	if err := sugar.Sync(); err != nil {
		log.Println(zapSyncErrorMessage)
	}
}

func Error(args ...any) {
	sugar.Errorln(args)
	if err := sugar.Sync(); err != nil {
		log.Println(zapSyncErrorMessage)
	}
}
