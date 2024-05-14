package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//DB is the global database connection
var DB *gorm.DB 

func ConnectDB(){
	//Read environment variables for database configuration
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

    customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})

	if err != nil {
		log.Fatalf("Failed to connect to database : %v", err)
	}

	DB = db 
	
	log.Println("Connected to database successfully")
}