package exchangeorderbook

type OrderBook struct {
	Bids SideOrders
	Asks SideOrders
}

func NewOrderBook(bids SideOrders, asks SideOrders) OrderBook {

	return OrderBook{
		Bids: bids,
		Asks: asks,
	}
}
