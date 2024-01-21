package database

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host, Port, Password, User, DBName, SSLMode string
}

func TestOpenConnection(t *testing.T) {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		"localhost", "5432", "riorusidano", "system-point", "disable",
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	assert.NotNil(t, db, "Database connection should not be nil")

	if err != nil {
		log.Fatal("Failed to get SQL database instance:", err)
	}
	sqlDB.Close()

}
