package usecase

import (
	"context"
	"suzushin54/event-sourcing-with-go/app/domain/aggregate"
	"suzushin54/event-sourcing-with-go/app/domain/command"
	"suzushin54/event-sourcing-with-go/app/domain/vo"
	"time"
)

type CreateInputV2 struct {
	OrderItems []*OrderItem `json:"items" validate:"required"`
	Contact    string       `json:"contact"`
}

type OrderItem struct {
	ItemName string `json:"item_name"`
	Price    uint32 `json:"price"`
	Quantity uint16 `json:"quantity"`
}

type CreateOutput struct {
	ID uint16
}

type UpdateStatusInput struct {
	ID     uint64 `json:"id"`
	Status uint16 `json:"status" validate:"required"`
}

func (o *orderUseCaseV2) CreateV2(ctx context.Context, in *CreateInputV2) (*CreateOutput, error) {
	orderID := uint64(o.snowflake.Generate().Int64())
	orderItems, err := o.convertInputToVo(in)
	if err != nil {
		return nil, err
	}

	// STEP: 注文作成コマンドを生成
	cmd, err := command.NewOrderCreateCommand(orderID, orderItems, in.Contact, time.Now().UTC())
	if err != nil {
		return nil, err
	}

	// STEP: 注文の集約を生成
	orderAggregate := aggregate.NewOrder(nil)

	// STEP: 集約にコマンドを渡して実行することで、集約の状態が変化する
	if err := orderAggregate.Execute(*cmd); err != nil {
		return nil, err
	}

	// STEP: 集約を永続化
	if err := o.orderRepositoryV2.Commit(ctx, orderAggregate); err != nil {
		return nil, err
	}

	return &CreateOutput{
		ID: uint16(cmd.ID),
	}, nil
}

func (o *orderUseCaseV2) convertInputToVo(in *CreateInputV2) ([]*vo.OrderItem, error) {
	// マスタ存在チェックなどは主題から外れるため割愛
	orderItems := make([]*vo.OrderItem, len(in.OrderItems))
	for i, v := range in.OrderItems {
		orderItem, err := vo.NewOrderItem(v.ItemName, v.Price, v.Quantity)
		if err != nil {
			return nil, err
		}
		orderItems[i] = orderItem
	}

	return orderItems, nil
}
