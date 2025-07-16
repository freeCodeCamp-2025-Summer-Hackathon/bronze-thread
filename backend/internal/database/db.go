package database

import (
	"os"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := os.Getenv("DB_URL")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = database

	// Automigrate all the models
	err = DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Item{},
		&models.SwapRequest{},
		&models.SwapRequestItem{},
		&models.SwapTransaction{},
		&models.SwapTransactionItem{},
		&models.ChatRoom{},
		&models.ChatMessage{},
		&models.UserRating{},
		&models.UserSession{},
	)
	return err
}
