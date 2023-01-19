package config

import "github.com/spf13/viper"

func ServerConfig() string {
	viper.SetDefault("SERVER_PORT", "8080")

	return ":" + viper.GetString("SERVER_PORT")
}
