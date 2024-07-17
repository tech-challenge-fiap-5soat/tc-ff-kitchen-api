package gateway

import (
	"fmt"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

type OrderGateway struct {
	Datasource interfaces.DatabaseSource
	OrderApi   interfaces.OrderApi
}

func NewOrderGateway(datasource interfaces.DatabaseSource,
	orderApi interfaces.OrderApi) interfaces.OrderGateway {
	return &OrderGateway{Datasource: datasource, OrderApi: orderApi}
}

func (og *OrderGateway) FindAll() ([]entity.Order, error) {
	orders, err := og.Datasource.FindAll("", "")

	if err != nil {
		return nil, err
	}

	foundOrders := []entity.Order{}

	for _, order := range orders {
		foundOrders = append(foundOrders, order.(entity.Order))
	}

	return foundOrders, nil
}

func (og *OrderGateway) FindById(id string) (*entity.Order, error) {
	order, err := og.Datasource.FindOne("_id", id)

	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	foundOrder := order.(*entity.Order)
	return foundOrder, nil
}

func (og *OrderGateway) FindAllByStatus(status valueobject.OrderStatus) ([]entity.Order, error) {
	orders, err := og.Datasource.FindAll("orderStatus", string(status))

	if err != nil {
		return nil, err
	}

	foundOrders := []entity.Order{}

	for _, order := range orders {
		foundOrders = append(foundOrders, order.(entity.Order))
	}

	return foundOrders, nil
}

func (og *OrderGateway) Save(order *entity.Order) (string, error) {
	insertResult, err := og.Datasource.Save(
		dto.OrderEntityToSaveRecordDTO(order),
	)

	if err != nil {
		return "", err
	}
	orderInserted := insertResult.(string)
	return orderInserted, nil
}

func (og *OrderGateway) Update(order *entity.Order) error {
	_, err := og.Datasource.Update(
		order.ID,
		dto.OrderEntityToUpdateRecordDTO(order),
	)

	if err != nil {
		return err
	}
	return nil
}

func (og *OrderGateway) ReleaseOrder(orderId string) error {
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

	err = og.OrderApi.ReleaseOrder(orderId)

	if err != nil {
		return err
	}
	return nil
}

func (og *OrderGateway) FinishOrder(orderId string) error {
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

	err = og.OrderApi.FinishOrder(orderId)

	if err != nil {
		return err
	}
	return nil
}
