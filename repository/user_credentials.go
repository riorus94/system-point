package repository

import (
	"context"
	"database/sql"
	"system-point/models/domain"
)

type UserCredentialsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, UserCredentials domain.UserCredentials) domain.UserCredentials
	Update(ctx context.Context, tx *sql.Tx, UserCredentials domain.UserCredentials) domain.UserCredentials
	Delete(ctx context.Context, tx *sql.Tx, UserCredentials domain.UserCredentials)
	FindById(ctx context.Context, tx *sql.Tx, UserCredentialID int) (domain.UserCredentials, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.UserCredentials
}
