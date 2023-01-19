package main

import (
	"fmt"
	"time"
	"tsbank/config"
	"tsbank/infra/database"
	"tsbank/migrations"
	"tsbank/routers"

	"github.com/spf13/viper"
)

func main() {
	// Timezone
	viper.SetDefault("TZ", "America/Sao_Paulo")
	location, _ := time.LoadLocation(viper.GetString("TZ"))
	time.Local = location

	if err := config.Config(); err != nil {
		fmt.Println("Configuration: ❌")
		panic(err)
	}
	fmt.Println("Configuration: ✅")

	// Database
	DSN := config.DbConfig()
	if err := database.DbConnection(DSN); err != nil {
		fmt.Println("Database: ❌")
		panic(err)
	}
	fmt.Println("Database: ✅")

	// Migrations
	migrations.Migrate()

	// Routes
	router := routers.SetupRoute()

	fmt.Println("Server: ✅")
	router.Run(config.ServerConfig())
}
