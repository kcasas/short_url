package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	logLevel, ok := os.LookupEnv("log_level")
	if !ok {
		logLevel = logrus.InfoLevel.String()
	}

	SetupLogger(logLevel)
}

// SetupLogger sets up global logrus settings
func SetupLogger(level string) {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	})

	logrus.SetOutput(os.Stdout)

	l, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Warnf("unknown log level: %q", level)
	} else {
		logrus.SetLevel(l)
	}
}
