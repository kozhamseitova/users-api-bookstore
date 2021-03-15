package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/tleuzhan13/bookstore/users-api/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	"log"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=$1"
	queryUpdateUser             = "UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4"
	queryDeleteUser             = "DELETE FROM users WHERE id=$1"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=$1"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=$1 AND password=$2 AND status=$3"
)

type UserRepository struct {
	Pool     *pgxpool.Pool
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func (r *UserRepository) Get(id int64) (*users.User, rest_errors.RestErr) {
	row := r.Pool.QueryRow(context.Background(), queryGetUser, id)

	var user *users.User
	user = new(users.User)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}
	return user, nil
}

func (r *UserRepository) Save(user *users.User) (*users.User, rest_errors.RestErr) {
	var userId int64
	row := r.Pool.QueryRow(context.Background(), queryInsertUser, user.FirstName, user.LastName, user.Email,
		user.DateCreated, user.Status, user.Password)
	err := row.Scan(&userId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}
	user.Id = userId
	return user, nil
}

func (r *UserRepository) Update(user *users.User) (*users.User, rest_errors.RestErr) {
	_, err := r.Pool.Exec(context.Background(), queryUpdateUser, user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}
	return user, nil
}

func (r *UserRepository) Delete(id int64) rest_errors.RestErr {
	_, err := r.Pool.Exec(context.Background(), queryDeleteUser, id)
	if err != nil {
		return rest_errors.NewInternalServerError("internal server error", err)
	}
	return nil
}

func (r *UserRepository) FindByStatus(status string) ([]*users.User, rest_errors.RestErr) {
	rows, err := r.Pool.Query(context.Background(), queryFindByStatus, status)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}

	results := make([]*users.User, 0)
	for rows.Next() {
		var user *users.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, rest_errors.NewInternalServerError("internal server error", err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (r *UserRepository) FindByEmailAndPassword(email string, password string) (*users.User, rest_errors.RestErr) {
	row := r.Pool.QueryRow(context.Background(), queryFindByEmailAndPassword, email, password, "active")

	var user *users.User
	user = new(users.User)

	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}
	return user, nil
}
