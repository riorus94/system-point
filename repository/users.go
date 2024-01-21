package repository

import (
	"context"
	"database/sql"
	"system-point/models/domain"
)

type UsersRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Users domain.Users) domain.Users
	Update(ctx context.Context, tx *sql.Tx, Users domain.Users) domain.Users
	Delete(ctx context.Context, tx *sql.Tx, Users domain.Users)
	FindById(ctx context.Context, tx *sql.Tx, UsersID int) (domain.Users, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Users
}
