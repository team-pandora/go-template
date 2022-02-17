package utils

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gobuffalo/envy"
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
	if _, ok := Truthy[envy.Get("TEST", "false")]; ok {
		logger.Out = ioutil.Discard
	}
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&formatter{})
	return logger
}
