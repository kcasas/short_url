package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	SetupLogger(viper.GetString("LOG_LEVEL"))
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
