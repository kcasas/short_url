package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("WEB_PORT", 8080)
	viper.SetDefault("LOG_LEVEL", logrus.InfoLevel.String())
}
