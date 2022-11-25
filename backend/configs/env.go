package configs

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

func MongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DB_URI")
}

func PORT() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PORT")
}