package repository_impl

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"suzushin54/event-sourcing-with-go/app/domain/aggregate"
	"suzushin54/event-sourcing-with-go/app/domain/event"
	"suzushin54/event-sourcing-with-go/app/domain/repository"
	"suzushin54/event-sourcing-with-go/app/infrastructure"
)

type OrderRepositoryV2 struct {
	eventStore infrastructure.EventStore
}

func NewOrderRepositoryV2(eventStore infrastructure.EventStore) repository.OrderRepositoryV2 {
	return &OrderRepositoryV2{
		eventStore: eventStore,
	}
}

func (o *OrderRepositoryV2) Find(ctx context.Context, id uint64) (*aggregate.Order, error) {
	// NOTE: 集約の復元を隠蔽したい場合はFindになると思われる
	// TODO: fetch未実装
	events, err := o.eventStore.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	order := aggregate.NewOrder(events)
	return order, nil
}

func (o *OrderRepositoryV2) LoadEvents(ctx context.Context, id uint64) ([]*event.DomainEvent, error) {
	events, err := o.eventStore.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (o *OrderRepositoryV2) Commit(ctx context.Context, order *aggregate.Order) error {
	if err := o.eventStore.Append(ctx, order.ID(), order.DomainEvents, aggregate.InitVersion); err != nil {
		return err
	}
	return nil
}

func (o *OrderRepositoryV2) CommitChanges(ctx context.Context, order *aggregate.Order, originalVersion uint32) error {
	// TODO: 元のバージョンをもらい、それ以降を永続化
	return nil

}
