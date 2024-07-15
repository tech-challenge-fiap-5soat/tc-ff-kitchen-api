package gateway

import (
	"fmt"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

type orderGateway struct {
	datasource interfaces.DatabaseSource
	orderApi   interfaces.OrderApi
}

func NewOrderGateway(datasource interfaces.DatabaseSource,
	orderApi interfaces.OrderApi) interfaces.OrderGateway {
	return &orderGateway{datasource: datasource, orderApi: orderApi}
}

func (og *orderGateway) FindAll() ([]entity.Order, error) {
	orders, err := og.datasource.FindAll("", "")

	if err != nil {
		return nil, err
	}

	foundOrders := []entity.Order{}

	for _, order := range orders {
		foundOrders = append(foundOrders, order.(entity.Order))
	}

	return foundOrders, nil
}

func (og *orderGateway) FindById(id string) (*entity.Order, error) {
	order, err := og.datasource.FindOne("_id", id)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	foundOrder := order.(*entity.Order)
	return foundOrder, nil
}

func (og *orderGateway) FindAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error) {
	orders, err := og.datasource.FindAll("orderStatus", string(status))

	if err != nil {
		return nil, err
	}

	foundOrders := []entity.Order{}

	for _, order := range orders {
		foundOrders = append(foundOrders, order.(entity.Order))
	}

	return foundOrders, nil
}

func (og *orderGateway) Save(order *entity.Order) (string, error) {
	insertResult, err := og.datasource.Save(
		dto.OrderEntityToSaveRecordDTO(order),
	)

	if err != nil {
		return "", err
	}
	orderInserted := insertResult.(string)
	return orderInserted, nil
}

func (og *orderGateway) Update(order *entity.Order) error {
	_, err := og.datasource.Update(
		order.ID,
		dto.OrderEntityToUpdateRecordDTO(order),
	)

	if err != nil {
		return err
	}
	return nil
}

func (og *orderGateway) ReleaseOrder(orderId string) error {
	order, err := og.FindById(orderId)

	if err != nil {
		return err
	}

	if order == nil {
		return fmt.Errorf("order not found")
	}

	if order.OrderStatus != valueobject.READY_TO_TAKEOUT {
		return fmt.Errorf("order cannot be released cause status is %s", order.OrderStatus.String())
	}

	err = og.orderApi.ReleaseOrder(orderId)

	if err != nil {
		return err
	}
	return nil
}

func (og *orderGateway) FinishOrder(orderId string) error {
	order, err := og.FindById(orderId)

	if err != nil {
		return err
	}

	if order == nil {
		return fmt.Errorf("order not found")
	}

	if order.OrderStatus != valueobject.COMPLETED {
		return fmt.Errorf("order cannot be finished cause status is %s", order.OrderStatus.String())
	}

	err = og.orderApi.FinishOrder(orderId)

	if err != nil {
		return err
	}
	return nil
}
