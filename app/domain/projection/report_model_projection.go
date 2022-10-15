package projection

import (
	"fmt"
	"suzushin54/event-sourcing-with-go/app/domain/event"
	"time"
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

// Order - レポートの文脈で使用する注文集約
type Order struct {
	id             uint64
	status         OrderStatus
	postponedCount uint16 // 提供時間を後ろに変更した回数
	version        uint32
	orderedAt      time.Time
	timeToServe    time.Time // 提供までにかかった時間
}

func NewOrder(events []event.DomainEvent) *Order {
	var o = &Order{}
	for _, e := range events {
		o.appendEvent(e)
	}
	return o
}

// appendEvent - イベントを受け取り、注文集約の状態を更新する
func (o *Order) appendEvent(e event.DomainEvent) {
	switch v := e.(type) {
	case event.OrderCreatedEvent:
		o.applyCreated(v)
	case event.OrderReceivedEvent:
		o.applyReceived(v)
	case event.PickupTimeChangedEvent:
		o.applyPickupTimeChanged(v)
	case event.OrderCancelledEvent:
		o.applyCancelled(v)
	default:
		fmt.Print("unknown event provided...")
	}
}

// applyCreated - 注文作成イベントを元に集約の状態を更新
func (o *Order) applyCreated(e event.OrderCreatedEvent) {
	o.id = e.OrderID
	o.status = OrderStatusOrderCreated
	o.orderedAt = e.OrderedAt
	o.version = InitVersion
}

// applyReceived - 店舗注文確認イベントを元に集約の状態を更新
func (o *Order) applyReceived(e event.OrderReceivedEvent) {
	o.status = OrderStatusOrderReceived
	o.version += 1
}

// applyPickupTimeChanged - 注文受け取り時間変更イベントを元に集約の状態を更新
func (o *Order) applyPickupTimeChanged(e event.PickupTimeChangedEvent) {
	o.postponedCount += 1
	o.version += 1
}

// applyCancelled - 注文キャンセルイベントを元に集約の状態を更新
func (o *Order) applyCancelled(e event.OrderCancelledEvent) {
	o.status = OrderStatusCancelled
	o.version += 1
}
