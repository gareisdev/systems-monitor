package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()

	// Set output format
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}
