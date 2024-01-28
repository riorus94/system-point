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

	kindActivityUsecase := usecase.NewKindActivityUsecase(databaseRepository)

	return Domain{
		KindActivityUsecase: &kindActivityUsecase,
	}
}
