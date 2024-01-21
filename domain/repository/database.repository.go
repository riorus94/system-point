package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DatabaseRepository interface {
	Create(value any) error
}

type databaseRepositoryImpl struct {
	database *gorm.DB
}

func NewDatabaseRepository(database *gorm.DB) databaseRepositoryImpl {
	return databaseRepositoryImpl{
		database: database,
	}
}

func (input databaseRepositoryImpl) Create(value any) error {
	result := input.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error create user:: %v", result.Error))
		return result.Error
	}

	return nil
}
