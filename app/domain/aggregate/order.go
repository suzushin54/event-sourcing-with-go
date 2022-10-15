package aggregate

import (
	"fmt"
	"suzushin54/event-sourcing-with-go/app/domain/command"
	"suzushin54/event-sourcing-with-go/app/domain/vo"
	"time"

	"suzushin54/event-sourcing-with-go/app/domain/event"
)

type OrderStatus uint16

const (
	OrderStatusNone OrderStatus = iota
	OrderStatusOrderCreated
	OrderStatusOrderReceived
	OrderStatusOrderReady
	OrderStatusCompleted
	OrderStatusCancelled
)

const InitVersion = 0

// Order - 注文集約
type Order struct {
	id                 uint64
	status             OrderStatus
	orderItems         []*vo.OrderItem
	version            uint32
	estimatePickupTime time.Time
	orderedAt          time.Time

	DomainEvents []*event.DomainEvent
}

// NewOrder - 受け取ったEventをReplayしてOrder集約を作成する
func NewOrder(events []*event.DomainEvent) *Order {
	var o = &Order{}
	for _, e := range events {
		o.appendEvent(*e)
	}
	return o
}

// appendEvent - イベントを受け取り、注文集約の状態を更新する
func (o *Order) appendEvent(e event.DomainEvent) {
	o.DomainEvents = append(o.DomainEvents, &e)

	switch v := e.(type) {
	case *event.OrderCreatedEvent:
		o.applyCreated(v)
	case *event.OrderReceivedEvent:
		o.applyReceived(v)
	case *event.PickupTimeChangedEvent:
		o.applyPickupTimeChanged(v)
	case *event.OrderCancelledEvent:
		o.applyCancelled(v)
	default:
		fmt.Printf("%T\n", v)
		fmt.Print("unknown event provided...")
	}
}

// applyCreated - 注文作成イベントを元に集約の状態を更新
func (o *Order) applyCreated(e *event.OrderCreatedEvent) {
	o.id = e.OrderID
	o.status = OrderStatusOrderCreated
	o.orderItems = e.OrderItems
	o.orderedAt = e.OrderedAt
	o.version = InitVersion
}

// applyReceived - 店舗注文確認イベントを元に集約の状態を更新
func (o *Order) applyReceived(e *event.OrderReceivedEvent) {
	o.status = OrderStatusOrderReceived
	o.version += 1
}

// applyPickupTimeChanged - 注文受け取り時間変更イベントを元に集約の状態を更新
func (o *Order) applyPickupTimeChanged(e *event.PickupTimeChangedEvent) {
	o.estimatePickupTime = e.ChangedTime
	o.version += 1
}

// applyCancelled - 注文キャンセルイベントを元に集約の状態を更新
func (o *Order) applyCancelled(e *event.OrderCancelledEvent) {
	o.status = OrderStatusCancelled
	o.version += 1
}

// NOTE: 他の言語であればOverloadingによりシグネチャが異なる同名メソッドを定義できる

// Execute - 任意のコマンドを受け取りイベントを作成し、集約に実行する
func (o *Order) Execute(c any) error {
	switch cmd := c.(type) {
	case command.OrderCreateCommand:
		orderCreatedEvent := event.NewOrderCreatedEvent(cmd.ID, cmd.OrderItems, cmd.Contact, cmd.OrderedAt)
		o.appendEvent(orderCreatedEvent)
	case command.OrderReceiveCommand:
		orderReceivedEvent := event.NewOrderReceivedEvent(cmd.OperatedBy, cmd.ReceivedAt)
		o.appendEvent(orderReceivedEvent)
	case command.OrderChangePickupTimeCommand:
		orderChangePickupTimeEvent := event.NewOrderPickupTimeChangedEvent(cmd.Cause, cmd.ChangedTime)
		o.appendEvent(orderChangePickupTimeEvent)
	case command.OrderCancelCommand:
		if o.status != OrderStatusOrderCreated {
			return fmt.Errorf("cannot cancel order. status: %d", o.status)
		}
		if o.orderedAt.Add(time.Hour).Before(cmd.CancelledAt) {
			return fmt.Errorf("order cannot be cancelled after 1 hour")
		}
		orderCancelledEvent := event.NewOrderCancelledEvent(cmd.Reason, cmd.CancelledAt)
		o.appendEvent(orderCancelledEvent)
	}
	return nil
}

// ExecuteCancel - キャンセル指示のみ切り出した例
func (o *Order) ExecuteCancel(cmd command.OrderCancelCommand) error {
	if o.status != OrderStatusOrderCreated {
		return fmt.Errorf("cannot cancel order. status: %d", o.status)
	}
	if o.orderedAt.Add(time.Hour).Before(cmd.CancelledAt) {
		return fmt.Errorf("order cannot be cancelled after 1 hour")
	}
	orderCancelledEvent := event.NewOrderCancelledEvent(cmd.Reason, cmd.CancelledAt)
	o.appendEvent(orderCancelledEvent)
	return nil
}

func (o *Order) ID() uint64 {
	return o.id
}

func (o *Order) Version() uint32 {
	return o.version
}
