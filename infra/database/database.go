package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

// DbConnection conecta com o banco de dados
func DbConnection(DSN string) error {
	var db = DB

	logMode := viper.GetBool("DB_LOGMODE")

	logLevel := logger.Silent
	if logMode {
		logLevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

// GetDB retorna a conexão com o banco de dados
func GetDB() *gorm.DB {
	return DB
}
