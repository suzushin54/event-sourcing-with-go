package command

import (
	"fmt"
	"suzushin54/event-sourcing-with-go/app/domain/vo"
	"time"
)

// OrderCreateCommand - 注文を受け取り、作成を指示するコマンド
type OrderCreateCommand struct {
	ID         uint64
	OrderItems []*vo.OrderItem
	Contact    string
	OrderedAt  time.Time
}

func NewOrderCreateCommand(id uint64, items []*vo.OrderItem, contact string, orderedAt time.Time) (*OrderCreateCommand, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("invalid arguments")
	}
	if contact == "" {
		return nil, fmt.Errorf("invalid arguments")
	}

	return &OrderCreateCommand{
		ID:         id,
		OrderItems: items,
		Contact:    contact,
		OrderedAt:  orderedAt,
	}, nil
}
