package infrastructure

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host, Port, Password, User, DBName, SSLMode string
}

func GetConnection() (*gorm.DB, error) {

	config := Config{
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		User:    os.Getenv("DB_USER"),
		DBName:  os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("SSL_MODE"),
	}
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil

}
