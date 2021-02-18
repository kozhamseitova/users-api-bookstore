package repository

import (
	"context"
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
	/*stmt, err := r.Pool.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	return nil*/
	return nil, nil
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
	/*stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
	}
	return nil*/
	return nil, nil
}

func (r *UserRepository) Delete(id int64) rest_errors.RestErr {
	/*	stmt, err := users_db.Client.Prepare(queryDeleteUser)
		if err != nil {
			logger.Error("error when trying to prepare delete user statement", err)
			return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
		}
		defer stmt.Close()

		if _, err = stmt.Exec(user.Id); err != nil {
			logger.Error("error when trying to delete user", err)
			return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
		}
		return nil*/
	return nil
}

func (r *UserRepository) FindByStatus(status string) ([]*users.User, rest_errors.RestErr) {
	/*stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
	}
	defer rows.Close()

	results := make([]users.User, 0)
	for rows.Next() {
		var user users.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, rest_errors.NewInternalServerError("error when tying to gett user", errors.New("database error"))
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil*/
	return nil, nil
}

func (r *UserRepository) FindByEmailAndPassword(email string, password string) (*users.User, rest_errors.RestErr) {
	/*stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerError("error when tying to find user", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, users.StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("error when tying to find user", errors.New("database error"))
	}
	return nil*/
	return nil, nil
}
