package lg

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	logFile, _ := os.OpenFile("n1ce.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Log = &logrus.Logger{
		Out:       logFile,
		Formatter: new(logrus.TextFormatter),
		Level:     logrus.DebugLevel,
	}
}
