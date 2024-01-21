package app

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host, Port, Password, User, DBName, SSLMode string
}

func GetConnection(config *Config) (*gorm.DB, error) {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s",
		"localhost", "5432", "riorusidano", "system-point", "disable",
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)

	// Use defer to ensure that the database connection is closed if there's an error during setup
	defer func() {
		if err != nil {
			_ = sqlDB.Close()
		}
	}()

	return db, nil

}
