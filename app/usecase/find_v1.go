package usecase

import (
	"context"
	"suzushin54/event-sourcing-with-go/app/domain/entity"
)

type FindOrderInputV1 struct {
	OrderID uint64
}

type FindOrderOutputV1 struct {
	Order *entity.Order
}

func (o *orderUseCaseV1) FindV1(ctx context.Context, in *FindOrderInputV1) (*FindOrderOutputV1, error) {
	order, err := o.repository.Find(ctx, in.OrderID)
	if err != nil {
		return nil, err
	}

	return &FindOrderOutputV1{
		Order: order,
	}, nil

}
