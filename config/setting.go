package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func initSetting() {
	viper.SetConfigName("settings")
	viper.AddConfigPath("$GOPATH/src/dictionary-of-chinese")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func init() {
	initSetting()
}
