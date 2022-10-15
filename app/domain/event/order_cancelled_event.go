package event

import (
	"encoding/json"
	"time"
)

const DomainEventNameOrderCancelled = "OrderCancelledEvent"

// OrderCancelledEvent - 注文キャンセルイベント
type OrderCancelledEvent struct {
	reason      string
	cancelledAt time.Time
}

func NewOrderCancelledEvent(reason string, cancelledAt time.Time) *OrderCancelledEvent {
	return &OrderCancelledEvent{reason: reason, cancelledAt: cancelledAt}
}

func (o OrderCancelledEvent) EventType() string {
	return DomainEventNameOrderCancelled
}

func (o OrderCancelledEvent) Reason() string {
	return o.reason
}

func (o OrderCancelledEvent) CancelledAt() time.Time {
	return o.cancelledAt
}

func (o OrderCancelledEvent) Serialize() (string, error) {
	bytes, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
