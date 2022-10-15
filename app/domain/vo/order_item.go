package vo

import (
	"fmt"
)

// OrderItem - ひとつの注文商品を表す値オブジェクト
type OrderItem struct {
	ItemName string `db:"item_name" json:"item_name"`
	Price    uint32 `db:"price" json:"price"`
	Quantity uint16 `db:"quantity" json:"quantity"`
}

func NewOrderItem(itemName string, price uint32, quantity uint16) (*OrderItem, error) {
	if itemName == "" || quantity == 0 {
		return nil, fmt.Errorf("invalid arguments")
	}

	return &OrderItem{
		ItemName: itemName,
		Price:    price,
		Quantity: quantity,
	}, nil
}

func (o OrderItem) Equals(v OrderItem) bool {
	return o.ItemName == v.ItemName && o.Price == v.Price && o.Quantity == v.Quantity
}
