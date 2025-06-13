package database

import (
	"log"
	"myclinic/internal/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func InitDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}else{
		log.Println("Connected to database")
	}

	if err := db.AutoMigrate(&models.User{}, &models.Patient{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
	return db
}
