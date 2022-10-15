package usecase

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/app/domain/entity"
	"suzushin54/event-sourcing-with-go/app/domain/vo"
	"suzushin54/event-sourcing-with-go/pkg"
	"time"
)

type CreateInputV1 struct {
	OrderItems []*OrderItemV1 `json:"items" validate:"required"`
	Contact    string         `json:"contact"`
}

type OrderItemV1 struct {
	ItemID   uint64 `json:"item_id"`
	ItemName string `json:"item_name"`
	Price    uint32 `json:"price"`
	Quantity uint16 `json:"quantity"`
}

type CreateOutputV1 struct {
	ID uint16
}

func (o *orderUseCaseV1) CreateV1(ctx context.Context, in *CreateInputV1) (*CreateOutputV1, error) {
	// STEP: Inputを元にEntityを生成
	orderEntity, err := o.convertInputToEntity(in, time.Now().UTC())
	if err != nil {
		return nil, err
	}

	if err := pkg.Transaction(
		ctx, o.conn, func(tx *sqlx.Tx) error {
			// STEP: Entityを永続化
			if err := o.repository.WithTx(tx).Save(ctx, orderEntity); err != nil {
				return err
			}
			return nil
		},
	); err != nil {
		return nil, fmt.Errorf("failed pkg.Transaction: %w", err)
	}

	return &CreateOutputV1{
		ID: uint16(orderEntity.ID),
	}, nil
}

func (o *orderUseCaseV1) convertInputToEntity(in *CreateInputV1, now time.Time) (*entity.Order, error) {
	orderItems := make([]*vo.OrderItem, len(in.OrderItems))
	for i, v := range in.OrderItems {
		orderItem, err := vo.NewOrderItem(v.ItemName, v.Price, v.Quantity)
		if err != nil {
			return nil, err
		}
		orderItems[i] = orderItem
	}

	orderID := o.snowflake.Generate().Int64()
	orderEntity, err := entity.NewOrder(uint64(orderID), orderItems, in.Contact, now)
	if err != nil {
		return nil, err
	}

	return orderEntity, nil
}
