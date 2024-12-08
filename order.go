package exchangeorderbook

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	id        int32
	side      OrderSide
	price     decimal.Decimal
	quantity  decimal.Decimal
	timestamp int64
}

func (o *Order) ID() int32 {

	return o.id
}

func (o *Order) Side() OrderSide {

	return o.side
}

func (o *Order) Quantity() decimal.Decimal {

	return o.quantity
}

func (o *Order) Price() decimal.Decimal {

	return o.price
}

func (o *Order) Timestamp() int64 {

	return o.timestamp
}

func (o *Order) Time() time.Time {

	return time.Time{}.Add(time.Duration(o.timestamp) * time.Millisecond)
}

func (o *Order) MarshalJSON() ([]byte, error) {

	return json.Marshal(
		&struct {
			ID        int32           `json:"id"`
			Side      OrderSide       `json:"side"`
			Price     decimal.Decimal `json:"price"`
			Quantity  decimal.Decimal `json:"quantity"`
			Timestamp time.Time       `json:"time"`
		}{
			ID:        o.ID(),
			Side:      o.Side(),
			Price:     o.Price(),
			Quantity:  o.Quantity(),
			Timestamp: o.Time(),
		},
	)
}

func (o *Order) UnmarshalJSON(data []byte) error {

	obj := struct {
		ID        int32           `json:"id"`
		Side      OrderSide       `json:"side"`
		Price     decimal.Decimal `json:"price"`
		Quantity  decimal.Decimal `json:"quantity"`
		Timestamp int64           `json:"time"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	o.id = obj.ID
	o.side = obj.Side
	o.price = obj.Price
	o.quantity = obj.Quantity
	o.timestamp = obj.Timestamp

	return nil
}

func NewOrder(id int32, side OrderSide, quantity decimal.Decimal, price decimal.Decimal, timestamp time.Time) *Order {

	return &Order{
		id:        id,
		side:      side,
		quantity:  quantity,
		price:     price,
		timestamp: timestamp.UnixMilli(),
	}
}
