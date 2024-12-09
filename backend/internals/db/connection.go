package db

import (
	"fmt"
	"log"

	"github.com/iacopoghilardi/mydget-backend/internals/config"
	"github.com/iacopoghilardi/mydget-backend/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	config := config.GetConfig()
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database: ", err)
		return err
	}
	DB = db
	log.Println("Connected to database")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func Ping() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func Close() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func Migrate() error {
	return DB.AutoMigrate(&models.User{})
}
