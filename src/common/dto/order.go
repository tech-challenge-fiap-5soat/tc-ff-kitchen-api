package dto

import (
	"time"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

type OrderCreateDTO struct {
	entity.Order
}

type OrderUpdateDTO struct {
	Cpf           string         `json:"customer"`
	OrderItemsDTO []OrderItemDTO `json:"orderItems"`
}

type OrderItemDTO struct {
	ProductId string `json:"product"`
	Quantity  int    `json:"quantity"`
}

func OrderEntityToSaveRecordDTO(order *entity.Order) map[string]interface{} {
	return map[string]interface{}{
		"_id":         order.ID,
		"customer":    order.Customer,
		"orderStatus": order.OrderStatus,
		"orderItems":  order.OrderItems,
		"amount":      order.Amount,
		"createdAt":   valueobject.CustomTime{Time: time.Now()},
	}
}

func OrderEntityToUpdateRecordDTO(order *entity.Order) map[string]interface{} {
	return map[string]interface{}{
		"orderStatus": order.OrderStatus,
		"orderItems":  order.OrderItems,
		"amount":      order.Amount,
		"updatedAt":   valueobject.CustomTime{Time: time.Now()},
	}
}
