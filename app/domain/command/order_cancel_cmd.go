package command

import "time"

// OrderCancelCommand - 注文キャンセルを指示するコマンド
type OrderCancelCommand struct {
	Reason      string
	CancelledAt time.Time
}

func NewOrderCancelCommand(
	reason string,
	cancelledAt time.Time,
) (*OrderCancelCommand, error) {
	return &OrderCancelCommand{
		Reason:      reason,
		CancelledAt: cancelledAt,
	}, nil
}
