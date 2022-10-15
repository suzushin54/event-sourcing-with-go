package command

import (
	"time"
)

// OrderReceiveCommand - 店舗注文確認を指示するコマンド
type OrderReceiveCommand struct {
	OperatedBy string
	ReceivedAt time.Time
}

func NewOrderReceiveCommand(operatedBy string, receivedAt time.Time) (*OrderReceiveCommand, error) {
	return &OrderReceiveCommand{
		OperatedBy: operatedBy,
		ReceivedAt: receivedAt,
	}, nil
}
