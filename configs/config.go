package configs

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigFile("configs/config.yml")

	return viper.ReadInConfig()
}
