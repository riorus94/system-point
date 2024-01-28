package usecase

import (
	"system-point/domain/entity"
	"system-point/domain/repository"
)

type KindActivityUsecase interface {
	SumAllPoints() (int, error)
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

func (input *kindActivityUsecaseImpl) SaveActivity(action entity.KindActivity) error {
	if err := input.databaseRepository.Create(&action); err != nil {
		return err
	}

	return nil
}

func (f *kindActivityUsecaseImpl) SumAllPoints() (int, error) {
	activities, err := f.databaseRepository.GetAllActivities()
	if err != nil {
		return 0, err
	}

	totalPoints := 0
	for _, activity := range activities {
		totalPoints += activity.Point
	}

	return totalPoints, nil
}
