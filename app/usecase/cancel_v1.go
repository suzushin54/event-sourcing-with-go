package usecase

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/pkg"
)

type CancelInputV1 struct {
	OrderID uint64
}

func (o *orderUseCaseV1) CancelV1(ctx context.Context, in *CancelInputV1) error {
	// STEP: Inputの注文IDを元にEntityを取得
	orderEntity, err := o.repository.Find(ctx, in.OrderID)
	if err != nil {
		return err
	}
	// STEP: Entityの振る舞いを呼び出して状態を変更
	if err := orderEntity.Cancel(time.Now()); err != nil {
		return err
	}
	if err := pkg.Transaction(
		ctx, o.conn, func(tx *sqlx.Tx) error {
			// STEP: 状態を変更させたEntityを永続化
			if err := o.repository.WithTx(tx).Update(ctx, orderEntity); err != nil {
				return err
			}
			return nil
		},
	); err != nil {
		return err
	}
	return nil
}
