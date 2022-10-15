package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/pkg"
)

type ReceiveInputV1 struct {
	OrderID uint64
}

func (o *orderUseCaseV1) ReceiveV1(ctx context.Context, in *ReceiveInputV1) error {
	orderEntity, err := o.repository.Find(ctx, in.OrderID)
	if err != nil {
		return err
	}

	if err := orderEntity.Receive(time.Now()); err != nil {
		return err
	}

	if err := pkg.Transaction(
		ctx, o.conn, func(tx *sqlx.Tx) error {
			repoTx := o.repository.WithTx(tx)
			if err := repoTx.Update(ctx, orderEntity); err != nil {
				return err
			}
			return nil
		},
	); err != nil {
		return fmt.Errorf("failed pkg.Transaction: %w", err)
	}
	return nil
}
