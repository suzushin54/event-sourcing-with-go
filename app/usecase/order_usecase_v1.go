package usecase

import (
	"context"
	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/app/domain/repository"
	"suzushin54/event-sourcing-with-go/pkg"
)

type OrderUseCaseV1 interface {
	CreateV1(ctx context.Context, in *CreateInputV1) (*CreateOutputV1, error)
	FindV1(ctx context.Context, in *FindOrderInputV1) (*FindOrderOutputV1, error)
	ReceiveV1(ctx context.Context, in *ReceiveInputV1) error
	CancelV1(ctx context.Context, in *CancelInputV1) error
}

type orderUseCaseV1 struct {
	conn       *sqlx.DB
	repository repository.OrderRepositoryV1
	snowflake  pkg.SnowflakeNode
}

func NewOrderUseCaseV1(
	conn *sqlx.DB,
	repo repository.OrderRepositoryV1,
	snowflake pkg.SnowflakeNode,
) OrderUseCaseV1 {
	return &orderUseCaseV1{
		conn:       conn,
		repository: repo,
		snowflake:  snowflake,
	}
}
