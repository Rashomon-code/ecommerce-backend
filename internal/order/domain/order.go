package domain

import (
	"errors"
	"time"
)

var (
	ErrOrderIDRequired    = errors.New("order ID is required")
	ErrCustomerIDRequired = errors.New("customer ID is required")
	ErrInvalidAmount      = errors.New("order amount must be greater than zero")

	ErrOrderAlreadyPaid   = errors.New("order is already paid")
	ErrCancelledOrder     = errors.New("cancelled order cannot be cancel")
	ErrInvalidOrderStatus = errors.New("invalid order status")
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

type Order struct {
	id          string
	customerID  string
	status      OrderStatus
	totalAmount int64
	createdAt   time.Time
}

func NewOrder(id, customerID string, totalAmount int64) (*Order, error) {
	if id == "" {
		return nil, ErrOrderIDRequired
	}
	if customerID == "" {
		return nil, ErrCustomerIDRequired
	}
	if totalAmount <= 0 {
		return nil, ErrInvalidAmount
	}

	return &Order{
		id:          id,
		customerID:  customerID,
		totalAmount: totalAmount,
		status:      OrderStatusPending,
		createdAt:   time.Now(),
	}, nil
}

func (o *Order) Pay() error {
	switch o.status {
	case OrderStatusPending:
		o.status = OrderStatusPaid
		return nil

	case OrderStatusPaid:
		return ErrOrderAlreadyPaid

	case OrderStatusCancelled:
		return ErrCancelledOrder

	default:
		return ErrInvalidOrderStatus
	}
}

func (o *Order) ID() string          { return o.id }
func (o *Order) CustomerID() string  { return o.customerID }
func (o *Order) TotalAmount() int64  { return o.totalAmount }
func (o *Order) Status() OrderStatus { return o.status }
