package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type formatter struct {
}

func (formatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelDesc := []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"}
	timestamp := entry.Time.Format("02/01/2006 15:04:05.000")
	return []byte(
		fmt.Sprintf("%s | [%s] %d %s %s (%s) %s\n",
			timestamp,
			levelDesc[entry.Level],
			entry.Data["status"],
			entry.Data["method"],
			entry.Data["path"],
			entry.Data["duration"],
			entry.Message,
		)), nil
}

// LoggerMiddleware handles logging all gin requests and errors.
func LoggerMiddleware() gin.HandlerFunc {
	logger := newLogger()

	return func(c *gin.Context) {
		// Request start time
		start := time.Now()

		// Process Request
		c.Next()

		// Set log info
		log := logger.WithFields(
			logrus.Fields{
				"status":   c.Writer.Status(),
				"method":   c.Request.Method,
				"path":     c.Request.URL.Path,
				"duration": time.Since(start).String(),
			},
		)

		message := "Request completed successfully."
		if err, exists := c.Get("error"); exists {
			message = err.(string)
		}

		switch {
		case isWarning(c):
			log.Warn(message)
		case isError(c):
			log.Error(message)
		default:
			log.Info(message)
		}
	}
}

// newLogger creates a new configured logrus logger.
func newLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&formatter{})
	return logger
}

// isWarning returns true if the response status in c is 4xx.
func isWarning(c *gin.Context) bool {
	return c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError
}

// isError returns true if the response status in c is >= 500.
func isError(c *gin.Context) bool {
	return c.Writer.Status() >= http.StatusInternalServerError
}
