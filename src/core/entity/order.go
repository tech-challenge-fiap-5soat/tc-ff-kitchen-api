package entity

import (
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

type Order struct {
	ID          string                  `json:"_id" bson:"_id"`
	Customer    Customer                `json:"customer,omitempty"`
	OrderStatus valueobject.OrderStatus `json:"orderStatus"`
	OrderItems  []OrderItem             `json:"orderItems"`
	CreatedAt   valueobject.CustomTime  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   valueobject.CustomTime  `json:"updatedAt" bson:"updatedAt"`
	Amount      float64                 `json:"amount"`
}

type OrderItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}
