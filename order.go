package exchangeorderbook

import (
	"github.com/shopspring/decimal"
)

type Order struct {
	id        int32
	side      OrderSide
	price     decimal.Decimal
	quantity  decimal.Decimal
	timestamp int64
}

func NewOrder(id int32, side OrderSide, price decimal.Decimal, quantity decimal.Decimal, timestamp int64) Order {

	return Order{
		id:        id,
		side:      side,
		price:     price,
		quantity:  quantity,
		timestamp: timestamp,
	}
}
