package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogging() *logrus.Logger {
	logger := logrus.New()
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// Configura el formato del log
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
