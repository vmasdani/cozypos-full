package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	// username := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbName)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(sqlite.Open("cozypos.sqlite3"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	sqlDB.SetMaxOpenConns(1)

	db.AutoMigrate(&ApiKey{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&StockIn{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&ItemTransaction{})

	dotenvErr := godotenv.Load()

	if dotenvErr != nil {
		log.Fatal("Error loading .env file.")

	}

	populateEnv := os.Getenv("POPULATE")

	if populateEnv == "YES" {
		populate(db)
	}

	return db
}
