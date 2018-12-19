package service

import (
	"os"
	"fmt"
	"github.com/my-stocks-pro/api-gateway-service/models"
	"github.com/jinzhu/gorm"
)

func (p *TypePSQL) MakeMigrations(connection *gorm.DB) {
	migrate := os.Getenv("MIGRATE")

	if migrate == "1" {
		fmt.Println("Migrate")

		connection.AutoMigrate(
			&models.Approve{},
			&models.Earnings{}, )

		//connection.AutoMigrate()

		fmt.Println("Migrations done")
	}
}
