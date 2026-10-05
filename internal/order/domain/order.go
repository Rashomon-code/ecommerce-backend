package domain

import (
	"errors"
	"time"
)

var (
	ErrOrderIDRequired    = errors.New("order ID is required")
	ErrCustomerIDRequired = errors.New("customer ID is required")
	ErrInvalidAmount      = errors.New("order amount must be greater than zero")
	ErrInvalidItem        = errors.New("order must contain at least one item")

	ErrOrderAlreadyPaid   = errors.New("order is already paid")
	ErrCancelledOrder     = errors.New("cancelled order cannot be pay")
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
	items       []OrderItem
	totalAmount int64
	createdAt   time.Time
}

type OrderItem struct {
	productID string
	price     int64
	quantity  int
}

func NewOrder(id, customerID string, items []OrderItem) (*Order, error) {
	if id == "" {
		return nil, ErrOrderIDRequired
	}
	if customerID == "" {
		return nil, ErrCustomerIDRequired
	}

	if len(items) == 0 {
		return nil, ErrInvalidItem
	}

	order := &Order{
		id:         id,
		customerID: customerID,
		status:     OrderStatusPending,
		items:      items,
		createdAt:  time.Now(),
	}

	order.calculateTotalAmount()

	return order, nil
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

func (o *Order) calculateTotalAmount() {
	var total int64
	for _, item := range o.items {
		total += item.Subtotal()
	}

	o.totalAmount = total
}

func (o *Order) ID() string          { return o.id }
func (o *Order) CustomerID() string  { return o.customerID }
func (o *Order) TotalAmount() int64  { return o.totalAmount }
func (o *Order) Status() OrderStatus { return o.status }

func (item OrderItem) Subtotal() int64 {
	return item.price * int64(item.quantity)
}
