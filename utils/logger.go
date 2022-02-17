package utils

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger = initLogger()

type formatter struct {
}

func (formatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelDesc := []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"}
	timestamp := entry.Time.Format("02/01/2006 15:04:05.000")
	return []byte(
		fmt.Sprintf("%s | [%s] %s\n",
			timestamp,
			levelDesc[entry.Level],
			entry.Message,
		)), nil
}

// LoggerMiddleware handles logging all gin requests and errors.
func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&formatter{})
	return logger
}
