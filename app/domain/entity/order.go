package entity

import (
	"fmt"
	"suzushin54/event-sourcing-with-go/app/domain/vo"
	"time"
)

const OrderPriceLimit = 10_000_000
const DefaultPickupTimeMinutes = 30

type OrderStatus uint16

const (
	OrderStatusNone OrderStatus = iota
	OrderStatusOrderCreate
	OrderStatusOrderReceive
	OrderStatusOrderReady
	OrderStatusOrderCompleted
	OrderStatusOrderCancelled
)

const InitVersion = 1

// Order - 一度の注文を表すモデル
type Order struct {
	ID                 uint64      `db:"id"`
	Status             OrderStatus `db:"status"`
	Items              []*vo.OrderItem
	Contact            string    `db:"contact"`
	Version            uint16    `db:"version"`
	EstimatePickupTime time.Time `db:"estimate_pickup_time"`
	OrderedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

func NewOrder(id uint64, items []*vo.OrderItem, contact string, createdAt time.Time) (*Order, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("invalid arguments")
	}
	if err := validateOrderPrice(items); err != nil {
		return nil, err
	}
	if contact == "" {
		return nil, fmt.Errorf("invalid arguments")
	}

	return &Order{
		ID:                 id,
		Status:             OrderStatusOrderCreate,
		Items:              items,
		Contact:            contact,
		Version:            InitVersion,
		EstimatePickupTime: createdAt.Add(time.Minute * DefaultPickupTimeMinutes),
		OrderedAt:          createdAt,
		UpdatedAt:          createdAt,
	}, nil
}

func validateOrderPrice(items []*vo.OrderItem) error {
	var totalPrice uint32
	for _, item := range items {
		totalPrice += item.Price
	}
	if totalPrice > OrderPriceLimit {
		return fmt.Errorf("exceeded amount available to order")
	}
	return nil
}

// RestoreOrder - 永続化データを元に復元する. infrastructure layer からのみ呼び出す.
func RestoreOrder(id uint64, status OrderStatus, items []*vo.OrderItem, contact string, version uint16, createdAt, updatedAt time.Time) *Order {
	return &Order{
		ID:        id,
		Status:    status,
		Items:     items,
		Contact:   contact,
		Version:   version,
		OrderedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (o *Order) Receive(updatedAt time.Time) error {
	if o.Status != OrderStatusOrderCreate {
		return fmt.Errorf("order cannot be changed to receive. current status is: %d", o.Status)
	}
	o.Status = OrderStatusOrderReceive
	o.UpdatedAt = updatedAt
	return nil
}

func (o *Order) UpdateEstimatePickupTime(estimatePickupTime time.Time, updatedAt time.Time) error {
	if o.Status != OrderStatusOrderReceive {
		return fmt.Errorf("order cannot be changed to modify estimate pickup time. current status is: %d", o.Status)
	}
	o.EstimatePickupTime = estimatePickupTime
	o.UpdatedAt = updatedAt
	return nil
}

func (o *Order) Cancel(updatedAt time.Time) error {
	if o.Status != OrderStatusOrderCreate {
		return fmt.Errorf("order cannot be changed to cancel. current status is: %d", o.Status)
	}
	if o.OrderedAt.Add(time.Hour).Before(updatedAt) {
		return fmt.Errorf("order cannot be changed to cancel. an hour has passed")
	}

	o.Status = OrderStatusOrderCancelled
	o.UpdatedAt = updatedAt
	return nil
}
