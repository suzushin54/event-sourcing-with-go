package infrastructure

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/app/domain/event"
)

type EventStore interface {
	Fetch(ctx context.Context, id uint64) ([]event.DomainEvent, error)
	Append(ctx context.Context, id uint64, newEvents []*event.DomainEvent, expectedVersion uint64) error
}

type MySqlEventStore struct {
	conn sqlx.ExtContext
}

func NewMySqlEventStore(conn *sqlx.DB) EventStore {
	return &MySqlEventStore{conn: conn}
}

func (m MySqlEventStore) Fetch(ctx context.Context, id uint64) ([]event.DomainEvent, error) {
	type OrderEventDao struct {
		OrderID   int    `db:"order_id"`
		Version   int    `db:"version"`
		EventType string `db:"event_type"`
		EventData string `db:"event_data"`
	}
	var events []*OrderEventDao
	q := `SELECT order_id, version, event_type, event_data FROM order_events WHERE order_id = ? ORDER BY version asc`

	err := sqlx.SelectContext(ctx, m.conn, &events, q, id)
	if err != nil {
		return nil, fmt.Errorf("failed sqlx.SelectContext: %w", err)
	}
	if len(events) == 0 {
		return nil, fmt.Errorf("failed to get order events. order_id: %d", id)
	}

	// TODO: 順序保証

	var domainEvents []event.DomainEvent
	for _, v := range events {
		switch v.EventType {
		case "OrderCreated":
			var e event.OrderCreatedEvent
			if err := e.Deserialize(v.EventData); err != nil {
				return nil, fmt.Errorf("failed to deserialize: %w", err)
			}
			domainEvents = append(domainEvents, e)
		case "OrderReceived":
			var e event.OrderReceivedEvent
			if err := e.Deserialize(v.EventData); err != nil {
				return nil, fmt.Errorf("failed to deserialize: %w", err)
			}
			domainEvents = append(domainEvents, e)
		case "OrderCanceled":
			var e event.OrderCancelledEvent
			if err := e.Deserialize(v.EventData); err != nil {
				return nil, fmt.Errorf("failed to deserialize: %w", err)
			}
			domainEvents = append(domainEvents, e)
		case "PickupTimeChanged":
			var e event.PickupTimeChangedEvent
			if err := e.Deserialize(v.EventData); err != nil {
				return nil, fmt.Errorf("failed to deserialize: %w", err)
			}
			domainEvents = append(domainEvents, e)
		}
	}
	return domainEvents, nil
}

func (m MySqlEventStore) Append(ctx context.Context, id uint64, newEvents []*event.DomainEvent, expectedVersion uint64) error {

	q := `INSERT INTO order_events (order_id, version, event_type, event_data) VALUES (:order_id, :version, :event_type, :event_data)`
	for i, v := range newEvents {
		jsonStr, err := (*v).Serialize()
		if err != nil {
			return err
		}
		if _, err := sqlx.NamedExecContext(
			ctx, m.conn, q, map[string]interface{}{
				"order_id":   id,
				"version":    expectedVersion + uint64(i+1), // NOTE: 0から始まるので+1
				"event_type": (*v).EventType(),
				"event_data": jsonStr,
			},
		); err != nil {
			return fmt.Errorf("failed sqlx.NamedExecContext: %w", err)
		}
	}

	return nil
}
