package event

import (
	"encoding/json"
	"suzushin54/event-sourcing-with-go/app/domain/vo"
	"time"
)

const DomainEventNameOrderCreated = "OrderCreated"

// OrderCreatedEvent - 注文作成イベント
type OrderCreatedEvent struct {
	OrderID    uint64          `json:"order_id"`
	OrderItems []*vo.OrderItem `json:"order_items"`
	Contact    string          `json:"contact"`
	OrderedAt  time.Time       `json:"ordered_at"`
}

func NewOrderCreatedEvent(orderID uint64, items []*vo.OrderItem, contact string, createdAt time.Time) *OrderCreatedEvent {
	return &OrderCreatedEvent{OrderID: orderID, OrderItems: items, Contact: contact, OrderedAt: createdAt}
}

func (o OrderCreatedEvent) EventType() string {
	return DomainEventNameOrderCreated
}

func (o OrderCreatedEvent) Serialize() (string, error) {
	bytes, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
