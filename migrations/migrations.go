package migrations

import (
	"tsbank/infra/database"
	"tsbank/models"
)

// Migrate executa as migrações
func Migrate() {
	var migrationModels = []interface{}{
		&models.User{},
		&models.Transaction{},
	}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
