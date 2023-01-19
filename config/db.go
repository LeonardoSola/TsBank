package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Driver   string
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
	SSLMode  string
	TimeZone string
}

// Database é a configuração do banco de dados
var Database DatabaseConfig

// DbConfig retorna a string de conexão com o banco de dados
func DbConfig() (DSN string) {
	Database = DatabaseConfig{
		Driver:   viper.GetString("DB_DRIVER"),
		DBName:   viper.GetString("DB_DBNAME"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
		TimeZone: viper.GetString("DB_TIMEZONE"),
	}

	DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		Database.Host,
		Database.User,
		Database.Password,
		Database.DBName,
		Database.Port,
		Database.SSLMode,
		Database.TimeZone,
	)

	return DSN
}
