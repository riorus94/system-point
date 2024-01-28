package domain

import (
	"system-point/domain/repository"
	"system-point/domain/usecase"
	"system-point/infrastructure"
)

type Domain struct {
	KindActivityUsecase usecase.KindActivityUsecase
}

func ConstructDomain() Domain {
	databaseConn, _ := infrastructure.GetConnection()

	databaseRepository := repository.NewDatabaseRepository(databaseConn)

	// registerUsecase := usecase.NewUserRegisterUsecase(databaseRepository)

	kindActivityUsecase := usecase.NewKindActivityUsecase(databaseRepository)

	return Domain{
		KindActivityUsecase: &kindActivityUsecase,
	}
}
