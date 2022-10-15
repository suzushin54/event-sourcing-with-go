package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/app/domain/entity"
)

type OrderRepositoryV1 interface {
	WithTx(tx *sqlx.Tx) OrderRepositoryV1
	Find(ctx context.Context, id uint64) (*entity.Order, error)
	Save(ctx context.Context, order *entity.Order) error
	Update(ctx context.Context, order *entity.Order) error
}
