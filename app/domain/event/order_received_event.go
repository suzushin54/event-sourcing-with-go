package event

import (
	"encoding/json"
	"time"
)

const DomainEventNameOrderReceived = "OrderReceived"

// OrderReceivedEvent - 店舗注文確認イベント
type OrderReceivedEvent struct {
	OperatedBy string
	ReceivedAt time.Time
}

func NewOrderReceivedEvent(operatedBy string, receivedAt time.Time) *OrderReceivedEvent {
	return &OrderReceivedEvent{
		OperatedBy: operatedBy,
		ReceivedAt: receivedAt,
	}
}

func (o OrderReceivedEvent) EventType() string {
	return DomainEventNameOrderReceived
}

func (o OrderReceivedEvent) Serialize() (string, error) {
	bytes, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (o OrderReceivedEvent) Deserialize(data string) error {
	if err := json.Unmarshal([]byte(data), &o); err != nil {
		return err
	}
	return nil
}
