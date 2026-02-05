package repository

import (
	"context"
	"database/sql"
	"fendi/modul-03-task/transport"
)

type CheckoutRepository struct {
	db *sql.DB
}

func NewCheckoutRepository(db *sql.DB) *CheckoutRepository {
	return &CheckoutRepository{db: db}
}

func (r *CheckoutRepository) CreateCheckoutTransaction(ctx context.Context, req transport.CheckoutRequest) error {
}
