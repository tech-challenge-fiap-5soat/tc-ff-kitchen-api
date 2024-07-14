package interfaces

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

type OrderUseCase interface {
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	GetAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error)
	CreateOrder(order entity.Order) (string, error)
	UpdateOrder(orderId string, order entity.Order) error
	UpdateOrderStatus(orderId string, status valueobject.OrderStatus) error
}

type OrderGateway interface {
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	FindAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error)
	Save(order *entity.Order) (string, error)
	Update(order *entity.Order) error
	FinishOrder(orderId string) error
}

type OrderController interface {
	FindAll() ([]entity.Order, error)
	FindById(id string) (*entity.Order, error)
	GetAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error)
	CreateOrder(order entity.Order) (string, error)
	UpdateOrder(orderId string, order entity.Order) error
	UpdateOrderStatus(orderId string, status valueobject.OrderStatus) error
}

type OrderApi interface {
	FinishOrder(orderId string) error
}
