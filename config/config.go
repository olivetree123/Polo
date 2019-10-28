package config

import (
	"github.com/spf13/viper"
	. "polo/common"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigFile("/etc/polo/config.toml")
	err := Config.ReadInConfig()
	if err != nil {
		Logger.Error(err)
	}
}
