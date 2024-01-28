package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DatabaseRepository interface {
	FindRandomActivity(data any, conditions ...any) error
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

func (input databaseRepositoryImpl) FindRandomActivity(data any, conditions ...any) error {
	result := input.database.Order("RANDOM()").Limit(1).Find(data, conditions...)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching book:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (input databaseRepositoryImpl) Create(value any) error {
	result := input.database.Create(value)

	if result.Error != nil {
		log.Printf(fmt.Sprintf("Error create user:: %v", result.Error))
		return result.Error
	}

	return nil
}
