package domain

import (
	"context"
	"errors"
)

var ErrOrderNotFound = errors.New("order not found")

type OrderRepository interface {
	Save(ctx context.Context, Order *Order) error
	FindByID(ctx context.Context, id string) (*Order, error)
}
