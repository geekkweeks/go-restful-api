package repository

import (
	"context"
	"database/sql"
	"geekkweeks/go-restful-api/model/domain"
)

type UserRepository interface {
	Add(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}
