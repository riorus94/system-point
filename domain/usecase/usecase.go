package usecase

import (
	"system-point/domain/entity"
	"system-point/domain/repository"
)

type KindActivityUsecase interface {
	GetRandomActivity() (*entity.KindActivity, error)
	SaveActivity(action entity.KindActivity) error
}

type kindActivityUsecaseImpl struct {
	databaseRepository repository.DatabaseRepository
}

func NewKindActivityUsecase(databaseRepository repository.DatabaseRepository) kindActivityUsecaseImpl {
	return kindActivityUsecaseImpl{
		databaseRepository: databaseRepository,
	}
}

func (input *kindActivityUsecaseImpl) GetRandomActivity() (*entity.KindActivity, error) {
	var activity entity.KindActivity
	if err := input.databaseRepository.FindRandomActivity(&activity); err != nil {
		return nil, err
	}

	return &activity, nil
}

func (input *kindActivityUsecaseImpl) SaveActivity(action entity.KindActivity) error {
	if err := input.databaseRepository.Create(&action); err != nil {
		return err
	}

	return nil
}
