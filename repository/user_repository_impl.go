package repository

import (
	"context"
	"database/sql"
	"errors"
	"geekkweeks/go-restful-api/helper"
	"geekkweeks/go-restful-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Add(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO USER(username, firstname, lastname, phone) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.FirstName, user.LastName, user.Phone)
	helper.PanicIfError(err)

	// Get last Id after inserted
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE USER set username = ?, firstName = ?, lastName  = ?, phone = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.FirstName, user.LastName, user.Phone)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "DELETE USER where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	SQL := "SELECT id, username, firstName, lastName, Phone from user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Phone)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User is not found")
	}
}

func (repository *UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, username, firstName, lastName, Phone from user"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Phone)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
