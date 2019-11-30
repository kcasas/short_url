package config

import (
	"math"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	DEFAULT_PREFIX_LENGTH = 2
	BASE                  = 62
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("WEB_PORT", 8080)
	viper.SetDefault("LOG_LEVEL", logrus.InfoLevel.String())

	viper.SetDefault("PREFIX_MIN", int64(math.Pow(BASE, DEFAULT_PREFIX_LENGTH-1)))
	viper.SetDefault("PREFIX_MAX", int64(math.Pow(BASE, DEFAULT_PREFIX_LENGTH)-1))
}
