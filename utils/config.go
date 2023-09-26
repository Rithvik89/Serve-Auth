package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	// Handle errors reading the config file

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetDB() string {
	return viper.GetString("DB_NAME")
}

func GetDBHost() string {
	return viper.GetString("DB_HOST")
}

func GetDBPort() int {
	return viper.GetInt("DB_PORT")
}

func GetDBUsername() string {
	return viper.GetString("DB_USERNAME")
}

func GetDBPassword() string {
	return viper.GetString("DB_PASSWORD")
}

func GetGithubClientID() string {
	return viper.GetString("GH_CLIENTID")
}

func GetGHClientSecret() string {
	return viper.GetString("GH_CLIENT_SECRET")
}
