package usecase

import (
	"fmt"
	"slices"
	"sort"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	orderStatus "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
)

type orderUseCase struct {
	gateway interfaces.OrderGateway
}

func NewOrderUseCase(gateway interfaces.OrderGateway) interfaces.OrderUseCase {
	return &orderUseCase{
		gateway: gateway,
	}
}

func (o *orderUseCase) FindAll() ([]entity.Order, error) {
	orders, err := o.gateway.FindAll()

	if err != nil {
		return nil, err
	}

	sort.Slice(orders, func(secondIndex, firstIndex int) bool {
		return sortByCreatedAt(orders[firstIndex], orders[secondIndex])
	})

	sort.Slice(orders, func(secondIndex, firstIndex int) bool {
		return sortByStatus(orders[firstIndex], orders[secondIndex])
	})

	var filtredOrders []entity.Order

	for _, order := range orders {
		if order.OrderStatus != orderStatus.COMPLETED {
			filtredOrders = append(filtredOrders, order)
		}
	}

	return filtredOrders, nil
}

func (o *orderUseCase) FindById(id string) (*entity.Order, error) {
	order, err := o.gateway.FindById(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *orderUseCase) GetAllByStatus(status orderStatus.OrderStatus) ([]entity.Order, error) {
	orders, err := o.gateway.FindAllByStatus(status)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderUseCase) CreateOrder(order entity.Order) (string, error) {

	orderToCreate := entity.Order{
		ID:          order.ID,
		OrderStatus: orderStatus.AWAITING_PREPARATION,
		OrderItems:  order.OrderItems,
		Amount:      order.Amount,
		Customer:    order.Customer,
	}

	orderId, err := o.gateway.Save(&orderToCreate)

	if err != nil {
		return "", err
	}

	return orderId, nil
}

func (o *orderUseCase) UpdateOrder(orderId string, order entity.Order) error {
	existentOrder, err := o.FindById(orderId)

	if err != nil {
		return err
	}

	if !existentOrder.OrderStatus.OrderCanBeUpdated() {
		return fmt.Errorf("order cannot be updated cause status is %s", existentOrder.OrderStatus.String())
	}

	orderToUpdate := entity.Order{
		ID:          orderId,
		OrderStatus: existentOrder.OrderStatus,
		OrderItems:  order.OrderItems,
		Amount:      order.Amount,
	}

	err = o.gateway.Update(&orderToUpdate)

	if err != nil {
		return err
	}

	return nil
}

func (o *orderUseCase) UpdateOrderStatus(orderId string, status orderStatus.OrderStatus) error {
	order, err := o.FindById(orderId)

	if err != nil {
		return err
	}

	if slices.Contains(order.OrderStatus.GetPreviousStatus(), status) {
		return fmt.Errorf(
			"order status %s cannot updated to previous status %s",
			order.OrderStatus.String(),
			status.String(),
		)
	}

	isValidNextStatus := order.OrderStatus.IsValidNextStatus(status.String())

	if !isValidNextStatus {
		return fmt.Errorf(
			"order status %s cannot be updated to %s. Status available are: %v",
			order.OrderStatus.String(),
			status.String(),
			order.OrderStatus.AvailableNextStatus(order.OrderStatus),
		)
	}

	err = o.updateOrderStatus(*order, status)

	if err != nil {
		return err
	}

	if order.OrderStatus.OrderIsReadyToTakeout(status) {
		o.gateway.ReleaseOrder(orderId)
	}

	if order.OrderStatus.OrderIsCompleted(status) {
		o.gateway.FinishOrder(orderId)
	}

	return nil
}

func (o *orderUseCase) updateOrderStatus(order entity.Order, newStatus orderStatus.OrderStatus) error {
	order.OrderStatus = newStatus
	return o.gateway.Update(&order)
}

func sortByStatus(firstOrder entity.Order, secondOrder entity.Order) bool {
	return (secondOrder.OrderStatus == orderStatus.READY_TO_TAKEOUT ||
		(secondOrder.OrderStatus == orderStatus.PREPARING && firstOrder.OrderStatus != orderStatus.READY_TO_TAKEOUT)) &&
		secondOrder.OrderStatus != firstOrder.OrderStatus
}

func sortByCreatedAt(firstOrder entity.Order, secondOrder entity.Order) bool {
	return !secondOrder.CreatedAt.Equal(firstOrder.CreatedAt.Time) &&
		secondOrder.CreatedAt.Before(firstOrder.CreatedAt.Time)
}
