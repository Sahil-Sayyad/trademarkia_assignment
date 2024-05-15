package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Sahil-Sayyad/Trademarkia/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database : %v", err)
	}

	DB = db 

	DB.AutoMigrate(&model.User{}, &model.Order{}, &model.Admin{}) 

	
	log.Println("Connected to database successfully")
}