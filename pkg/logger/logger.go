package logger

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

var baseLoggerTextFormaterConfig = &logrus.TextFormatter{
	FullTimestamp: true,
	ForceColors:   true,
	PadLevelText:  true,
}

func InitLogger() *logrus.Logger {
	log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: baseLoggerTextFormaterConfig,
		Level:     logrus.DebugLevel,
	}

	return log
}

func Get() *logrus.Logger {
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
