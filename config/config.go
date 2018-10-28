package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type GeneralConfig struct {
	ETCDHosts []string
	LogLevel  string
}

var config *GeneralConfig

func LoadConfig(filepath string) (a *GeneralConfig) {
	viper.SetConfigFile(filepath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	newConfig()
	return config
}

func LoadRemoteConfig() (a *GeneralConfig) {

	viper.AddRemoteProvider("etcd", "http://127.0.0.1:2379", "us_dev/config.yml")
	viper.SetConfigType("yml")
	err := viper.ReadRemoteConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	newConfig()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			err := viper.WatchRemoteConfig()
			if err != nil {
				fmt.Errorf("unable to read remote config: %v", err)
				continue
			}
			newConfig()
		}
	}()

	return config
}

func newConfig() *GeneralConfig {
	if config == nil {
		config = new(GeneralConfig)
	}
	config.ETCDHosts = viper.GetStringSlice("etcd_hosts")
	config.LogLevel = viper.GetString("log_level")
	l, _ := logrus.ParseLevel(config.LogLevel)
	logrus.SetLevel(l)
	return config
}
