package database

import (
	"fmt"
	"log"
	"to-do-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres port=5432 sslmode=disable Timezone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatal("Failed to migrate model")
	}

	DB = db
	fmt.Println("Database Connected Successfully")
}
