package command

import (
	"time"
)

// OrderChangePickupTimeCommand - 注文受取時間の変更を指示するコマンド
type OrderChangePickupTimeCommand struct {
	Cause       string
	ChangedTime time.Time
}

func NewOrderChangePickupTimeCommand(cause string, changedTime time.Time) (*OrderChangePickupTimeCommand, error) {
	return &OrderChangePickupTimeCommand{
		Cause:       cause,
		ChangedTime: changedTime,
	}, nil
}
