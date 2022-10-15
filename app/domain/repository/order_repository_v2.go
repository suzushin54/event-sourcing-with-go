package repository

import (
	"context"
	"suzushin54/event-sourcing-with-go/app/domain/aggregate"
	"suzushin54/event-sourcing-with-go/app/domain/event"
)

type OrderRepositoryV2 interface {
	Find(ctx context.Context, id uint64) (*aggregate.Order, error)
	LoadEvents(ctx context.Context, id uint64) ([]*event.DomainEvent, error)
	Commit(ctx context.Context, order *aggregate.Order) error
	CommitChanges(ctx context.Context, order *aggregate.Order, originalVersion uint32) error
}
