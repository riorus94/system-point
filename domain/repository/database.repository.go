package repository

import (
	"system-point/domain/entity"

	"gorm.io/gorm"
)

type DatabaseRepository interface {
	GetAllActivities() ([]entity.KindActivity, error)
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

func (r databaseRepositoryImpl) GetAllActivities() ([]entity.KindActivity, error) {
	var activities []entity.KindActivity
	if err := r.database.Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (input databaseRepositoryImpl) Create(value any) error {
	result := input.database.Create(value)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
