package db

import (
	"fmt"
	"log"
	"os"

	"github.com/RhnAdi/Gomle/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("message: failed to read env \n error: %F", err)
	}
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println("Success Connect Database Error => ", err)
	}

	// Migration Database
	db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Post{},
		&models.Image{},
		&models.Comment{},
	)

	return db, err
}

func DBClose(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		log.Panicf("DB close failed.")
	}
	database.Close()
}
