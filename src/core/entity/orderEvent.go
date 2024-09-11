package entity

type OrderEvent struct {
	Id          string `json:"id"`
	EventType   string `json:"eventType"`
	OrderStatus string `json:"orderStatus"`
	Order       *Order `json:"order"`
}
