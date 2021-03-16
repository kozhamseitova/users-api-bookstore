package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	users "gitlab.com/tleuzhan13/bookstore/users-api/domain"
	"gitlab.com/tleuzhan13/bookstore/users-api/foundation/rest_errors"
	"log"
)

const (
	MakeOrder = "INSERT INTO order(id, user_id, book_id) VALUES ($1,$2,$3) RETURNING id"
	Buy       = "DELETE FROM order WHERE id=$1"
)

type OrderRepo struct {
	Pool     *pgxpool.Pool
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func (r *OrderRepo) MakeOrder(order *users.Order) (*users.Order, rest_errors.RestErr) {
	var orderId int64
	row := r.Pool.QueryRow(context.Background(), MakeOrder, order.ID, order.UserID, order.BookID)
	err := row.Scan(&orderId)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal server error", err)
	}
	order.ID = orderId
	return order, nil
}

func (r *OrderRepo) Buy(id int64) rest_errors.RestErr {
	_, err := r.Pool.Exec(context.Background(), Buy, id)
	if err != nil {
		return rest_errors.NewInternalServerError("internal server error", err)
	}
	return nil
}
