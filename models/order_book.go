package models

// создаем структуру ордера
type DepthOrder struct {
	Price   float64 `json:"Price"`
	BaseQty float64 `json:"base_qty"`
}

// создаем книгу
type OrderBook struct {
	ID       int64
	Exchange string
	Pair     string
	Asks     []DepthOrder
	Bids     []DepthOrder
}
