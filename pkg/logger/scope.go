package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var scopeTagWidth uint = 10

const (
	resetColor = "\033[0m"
	cyan       = "\033[36m"
)

type scopeTextFormatter struct {
	scopeStr     string
	scopeStrLen  uint
	textFormater *logrus.TextFormatter
}

func (formater scopeTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if formater.scopeStrLen == 0 {
		return formater.textFormater.Format(entry)
	}

	entry.Message = fmt.Sprintf("%-*s %s", scopeTagWidth, formater.scopeStr, entry.Message)
	return formater.textFormater.Format(entry)
}

func GetWithScopes(scope ...string) *logrus.Logger {
	scopeStr, scopeStrLen := createScopesString(scope)
	if scopeStrLen > scopeTagWidth {
		scopeTagWidth = scopeStrLen
	}

	return &logrus.Logger{
		Out: os.Stdout,
		Formatter: scopeTextFormatter{
			scopeStr:     scopeStr,
			scopeStrLen:  scopeStrLen,
			textFormater: baseLoggerTextFormaterConfig,
		},
		Level: logrus.DebugLevel,
	}
}

func createScopesString(scopes []string) (string, uint) {
	if len(scopes) == 0 {
		return "", 0
	}
	str := colorTextCyan(fmt.Sprint("[", strings.Join(scopes, "]["), "]"))
	return str, uint(len([]rune(str)))
}

func colorTextCyan(str string) string {
	return cyan + str + resetColor
}
