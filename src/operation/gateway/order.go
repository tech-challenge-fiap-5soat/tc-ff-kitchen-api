package gateway

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/dto"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

type OrderGateway struct {
	Datasource       interfaces.DatabaseSource
	OrderApi         interfaces.OrderApi
	PublisherGateway interfaces.PublisherGateway
}

func NewOrderGateway(datasource interfaces.DatabaseSource,
	orderApi interfaces.OrderApi, publisherGateway interfaces.PublisherGateway) interfaces.OrderGateway {
	return &OrderGateway{Datasource: datasource, OrderApi: orderApi, PublisherGateway: publisherGateway}
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

	err = og.PublishEvent(order)

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

	err = og.PublishEvent(order)

	if err != nil {
		return err
	}
	return nil
}

func (og *OrderGateway) PublishEvent(order *entity.Order) error {

	event := entity.OrderEvent{
		Id:          uuid.New().String(),
		EventType:   order.OrderStatus.String(),
		OrderStatus: string(order.OrderStatus),
		Order:       order,
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error occurred while encoding order data: %s", err.Error())
	}

	queueName := og.PublisherGateway.GetQueueUrl()
	data := string(jsonData)
	err = og.PublisherGateway.PublishMessage(queueName, data)
	if err != nil {
		return err
	}
	return nil
}
