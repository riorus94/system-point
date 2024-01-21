package domain

import (
	"system-point/domain/repository"
	"system-point/domain/usecase"
	"system-point/infrastructure"
)

type Domain struct {
	UserRegistrationUsecase usecase.UserRegiserterUsecase
}

func ConstructDomain() Domain {
	databaseConn, _ := infrastructure.GetConnection()

	databaseRepository := repository.NewDatabaseRepository(databaseConn)

	registerUsecase := usecase.NewUserRegisterUsecase(databaseRepository)

	return Domain{
		UserRegistrationUsecase: &registerUsecase,
	}
}
