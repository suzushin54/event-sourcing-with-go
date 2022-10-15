package usecase

import (
	"context"
	"suzushin54/event-sourcing-with-go/app/domain/repository"
	"suzushin54/event-sourcing-with-go/pkg"
)

type OrderUseCaseV2 interface {
	CreateV2(ctx context.Context, in *CreateInputV2) (*CreateOutput, error)
	CancelV2(ctx context.Context, in *CancelInputV2) error
}

type orderUseCaseV2 struct {
	orderRepositoryV2 repository.OrderRepositoryV2
	snowflake         pkg.SnowflakeNode
}

func NewOrderUseCaseV2(
	orderRepositoryV2 repository.OrderRepositoryV2,
	snowflake pkg.SnowflakeNode,
) OrderUseCaseV2 {
	return &orderUseCaseV2{
		orderRepositoryV2: orderRepositoryV2,
		snowflake:         snowflake,
	}
}
