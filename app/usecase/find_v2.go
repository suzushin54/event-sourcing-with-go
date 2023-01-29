package usecase

import (
	"context"
	"suzushin54/event-sourcing-with-go/app/domain/aggregate"
)

type FindOrderInputV2 struct {
	OrderID uint64
}

type FindOrderOutputV2 struct {
	Order *aggregate.Order
}

func (o *orderUseCaseV2) FindV2(ctx context.Context, in *FindOrderInputV2) (*FindOrderOutputV2, error) {
	order, err := o.orderRepositoryV2.Find(ctx, in.OrderID)
	if err != nil {
		return nil, err
	}

	return &FindOrderOutputV2{
		Order: order,
	}, nil

}
