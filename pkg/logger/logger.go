package logger

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() *logrus.Logger {
	log = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
			ForceColors:   true,
			PadLevelText:  true,
		},
		Level: logrus.DebugLevel,
	}

	return log
}

func GetLogger() *logrus.Logger {
	return log
}

func PrintStruct(strct any, messages ...string) {
	jsonData, err := json.MarshalIndent(strct, "", "  ")
	if err != nil {
		log.Errorf("Failed to print given struct. Error: %v", err)
		return
	}

	log.Infof("%s %s", strings.Join(messages, " "), string(jsonData))
}
