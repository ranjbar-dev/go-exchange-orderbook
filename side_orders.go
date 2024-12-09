package exchangeorderbook

import (
	"github.com/shopspring/decimal"
)

type SideOrders struct {
	orders      map[int32]Order
	volume      decimal.Decimal
	totalOrders int
	depth       int
}

func NewSideOrders(orders map[int32]Order, volume decimal.Decimal, totalOrders int, depth int) SideOrders {

	return SideOrders{
		orders:      orders,
		volume:      volume,
		totalOrders: totalOrders,
		depth:       depth,
	}
}
