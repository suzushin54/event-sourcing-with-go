package repository_impl

import (
	"context"
	"fmt"
	"suzushin54/event-sourcing-with-go/app/domain/vo"

	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/app/domain/entity"
	"suzushin54/event-sourcing-with-go/app/domain/repository"
)

type OrderRepositoryV1 struct {
	conn sqlx.ExtContext
}

func NewOrderRepositoryV1(conn *sqlx.DB) repository.OrderRepositoryV1 {
	return &OrderRepositoryV1{
		conn: conn,
	}
}

func (o *OrderRepositoryV1) WithTx(tx *sqlx.Tx) repository.OrderRepositoryV1 {
	return &OrderRepositoryV1{
		conn: tx,
	}
}

func (o *OrderRepositoryV1) Find(ctx context.Context, id uint64) (*entity.Order, error) {
	var items []*vo.OrderItem
	q1 := `SELECT item_name, price, quantity FROM order_items WHERE order_id = ?`
	if err := sqlx.SelectContext(ctx, o.conn, &items, q1, id); err != nil {
		return nil, fmt.Errorf("failed sqlx.SelectContext: %w", err)
	}
	if len(items) == 0 {
		return nil, fmt.Errorf("failed to get order items. order_id: %d", id)
	}

	var order entity.Order
	q2 := `SELECT id, status, contact, version, created_at, updated_at FROM orders WHERE id = ?`
	if err := sqlx.GetContext(ctx, o.conn, &order, q2, id); err != nil {
		return nil, fmt.Errorf("failed sqlx.SelectContext: %w", err)
	}

	orderEntity := entity.RestoreOrder(
		order.ID,
		order.Status,
		items,
		order.Contact,
		order.Version,
		order.OrderedAt,
		order.UpdatedAt,
	)
	return orderEntity, nil
}

func (o *OrderRepositoryV1) Save(ctx context.Context, order *entity.Order) error {
	// Ordersテーブルに注文を永続化
	q1 := `INSERT INTO orders (id, status, contact, version, created_at, updated_at) VALUES (:id, :status, :contact, :version, :created_at, :updated_at)`
	if _, err := sqlx.NamedExecContext(
		ctx, o.conn, q1, map[string]interface{}{
			"id":         order.ID,
			"status":     order.Status,
			"contact":    order.Contact,
			"version":    order.Version,
			"created_at": order.OrderedAt,
			"updated_at": order.UpdatedAt,
		},
	); err != nil {
		return fmt.Errorf("failed sqlx.NamedExecContext: %w", err)
	}

	// OrderItemsテーブルに注文商品を永続化
	q2 := `INSERT INTO order_items (order_id, item_name, price, quantity)
		VALUES (:order_id, :item_name, :price, :quantity)`
	for _, item := range order.Items {
		// 必要に応じてBulk insertに
		if _, err := sqlx.NamedExecContext(
			ctx, o.conn, q2, map[string]interface{}{
				"order_id":  order.ID, // ここで集約のIDが必要
				"item_name": item.ItemName,
				"price":     item.Price,
				"quantity":  item.Quantity,
			},
		); err != nil {
			return fmt.Errorf("failed sqlx.NamedExecContext: %w", err)
		}
	}

	return nil
}

func (o *OrderRepositoryV1) Update(ctx context.Context, order *entity.Order) error {
	// Ordersテーブルの注文を更新
	q := `UPDATE orders SET status = :status, contact = :contact, version = :version, updated_at = :updated_at
			WHERE id = :id AND version = :original_version`

	if _, err := sqlx.NamedExecContext(
		ctx, o.conn, q, map[string]interface{}{
			"status":           order.Status,
			"contact":          order.Contact,
			"version":          order.Version + 1,
			"updated_at":       order.UpdatedAt,
			"id":               order.ID,
			"original_version": order.Version,
		},
	); err != nil {
		return fmt.Errorf("failed sqlx.NamedExecContext: %w", err)
	}

	// NOTE: 今回はOrderItemsは変更されないもの仮定して、order_itemsテーブルの更新処理を省略する

	return nil
}
