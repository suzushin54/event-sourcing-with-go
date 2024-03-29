package event

import (
	"encoding/json"
	"time"
)

const DomainEventNamePickupTimeChanged = "PickupTimeChanged"

// PickupTimeChangedEvent - 注文受取予定時間変更イベント
type PickupTimeChangedEvent struct {
	Cause       string
	ChangedTime time.Time
}

func NewOrderPickupTimeChangedEvent(cause string, changedTime time.Time) *PickupTimeChangedEvent {
	return &PickupTimeChangedEvent{
		Cause:       cause,
		ChangedTime: changedTime,
	}
}

func (o PickupTimeChangedEvent) EventType() string {
	return DomainEventNamePickupTimeChanged
}

func (o PickupTimeChangedEvent) Serialize() (string, error) {
	bytes, err := json.Marshal(o)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (o PickupTimeChangedEvent) Deserialize(data string) error {
	if err := json.Unmarshal([]byte(data), &o); err != nil {
		return err
	}
	return nil
}
