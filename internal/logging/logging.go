package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init(level string) {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.WarnLevel)
	}

	logrus.SetLevel(logLevel)
}

func SetLogLevel(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.WarnLevel)
		return
	}
	logrus.SetLevel(logLevel)
}
