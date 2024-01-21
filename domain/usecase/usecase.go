package usecase

import (
	"system-point/domain/entity"
	"system-point/domain/repository"
)

type UserRegiserterUsecase interface {
	Register(action entity.Register) error
}

type userRegisterImpl struct {
	databaseRepository repository.DatabaseRepository
}

func NewUserRegisterUsecase(databaseRepository repository.DatabaseRepository) userRegisterImpl {
	return userRegisterImpl{
		databaseRepository: databaseRepository,
	}
}

func (input *userRegisterImpl) Register(action entity.Register) error {
	if err := input.databaseRepository.Create(&action); err != nil {
		return err
	}

	return nil
}
