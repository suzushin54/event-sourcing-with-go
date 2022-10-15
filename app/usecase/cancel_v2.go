package usecase

import (
	"context"
	"suzushin54/event-sourcing-with-go/app/domain/aggregate"
	"suzushin54/event-sourcing-with-go/app/domain/command"
	"time"
)

type CancelInputV2 struct {
	OrderID uint64
	Reason  string
}

func (o *orderUseCaseV2) CancelV2(ctx context.Context, in *CancelInputV2) error {
	// STEP: イベントをロード
	events, err := o.orderRepositoryV2.LoadEvents(ctx, in.OrderID)
	if err != nil {
		return err
	}
	// STEP: イベントを元に集約を復元（RepositoryがAggregateを直接返す実装にするならここは不要）
	orderAggregate := aggregate.NewOrder(events)
	originalVersion := orderAggregate.Version()
	// STEP: 注文キャンセルコマンドを生成
	cmd, err := command.NewOrderCancelCommand(in.Reason, time.Now().UTC())
	if err != nil {
		return err
	}
	// STEP: 集約にコマンドを渡して実行することで、集約の状態が変化する
	if err := orderAggregate.Execute(*cmd); err != nil {
		return err
	}
	// STEP: 集約を永続化
	if err := o.orderRepositoryV2.CommitChanges(ctx, orderAggregate, originalVersion); err != nil {
		return err
	}
	return nil
}
