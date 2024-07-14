package controller

import (
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/usecase"
	vo "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/operation/gateway"
)

type OrderController struct {
	useCase interfaces.OrderUseCase
}

func NewOrderController(datasource interfaces.DatabaseSource, orderApi interfaces.OrderApi) interfaces.OrderController {

	gateway := gateway.NewOrderGateway(datasource, orderApi)
	return &OrderController{
		useCase: usecase.NewOrderUseCase(gateway),
	}
}

func (oc *OrderController) FindAll() ([]entity.Order, error) {
	return oc.useCase.FindAll()
}

func (oc *OrderController) FindById(id string) (*entity.Order, error) {
	return oc.useCase.FindById(id)
}

func (oc *OrderController) GetAllByStatus(status vo.OrderStatus) ([]entity.Order, error) {
	return oc.useCase.GetAllByStatus(status)
}

func (oc *OrderController) CreateOrder(order entity.Order) (string, error) {
	return oc.useCase.CreateOrder(order)
}

func (oc *OrderController) UpdateOrder(orderId string, order entity.Order) error {
	return oc.useCase.UpdateOrder(orderId, order)
}

func (oc *OrderController) UpdateOrderStatus(orderId string, status vo.OrderStatus) error {
	return oc.useCase.UpdateOrderStatus(orderId, status)
}
