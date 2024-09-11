package gateway_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/operation/gateway"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/test/mocks"
)

type MockDatasource struct {
	mock.Mock
}

func (m *MockDatasource) FindAll(arg1, arg2 string) ([]interface{}, error) {
	args := m.Called(arg1, arg2)
	return args.Get(0).([]interface{}), args.Error(1)
}

func TestOrderGateway_FindAll(t *testing.T) {
	mockDatasource := mocks.NewMockDatabaseSource(t)
	expectedOrders := []interface{}{
		entity.Order{ID: "1", Amount: 10},
		entity.Order{ID: "2", Amount: 20},
	}
	mockDatasource.On("FindAll", "", "").Return(expectedOrders, nil)

	og := &gateway.OrderGateway{Datasource: mockDatasource}

	orders, err := og.FindAll()

	assert.NoError(t, err)
	assert.Len(t, orders, 2)
	assert.Equal(t, "1", orders[0].ID)
	assert.Equal(t, 10.0, orders[0].Amount)
	assert.Equal(t, "2", orders[1].ID)
	assert.Equal(t, 20.0, orders[1].Amount)

	mockDatasource.AssertExpectations(t)
}
func TestOrderGateway_FindById(t *testing.T) {
	mockDatasource := mocks.NewMockDatabaseSource(t)
	expectedOrder := &entity.Order{ID: "1", Amount: 10}
	mockDatasource.On("FindOne", "_id", "1").Return(expectedOrder, nil)

	og := &gateway.OrderGateway{Datasource: mockDatasource}

	order, err := og.FindById("1")

	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "1", order.ID)
	assert.Equal(t, 10.0, order.Amount)

	mockDatasource.AssertExpectations(t)
}
func TestOrderGateway_FindAllByStatus(t *testing.T) {
	mockDatasource := mocks.NewMockDatabaseSource(t)
	expectedOrders := []interface{}{
		entity.Order{ID: "1", Amount: 10, OrderStatus: valueobject.OrderStatus("READY_TO_TAKEOUT")},
		entity.Order{ID: "2", Amount: 20, OrderStatus: valueobject.OrderStatus("READY_TO_TAKEOUT")},
	}
	mockDatasource.On("FindAll", "orderStatus", "READY_TO_TAKEOUT").Return(expectedOrders, nil)

	og := &gateway.OrderGateway{Datasource: mockDatasource}

	orders, err := og.FindAllByStatus(valueobject.OrderStatus("READY_TO_TAKEOUT"))

	assert.NoError(t, err)
	assert.Len(t, orders, 2)
	assert.Equal(t, "1", orders[0].ID)
	assert.Equal(t, 10.0, orders[0].Amount)
	assert.Equal(t, "2", orders[1].ID)
	assert.Equal(t, 20.0, orders[1].Amount)

	mockDatasource.AssertExpectations(t)
}
func TestOrderGateway_ReleaseOrder(t *testing.T) {
	mockDatasource := mocks.NewMockDatabaseSource(t)
	mockOrderApi := mocks.NewMockOrderApi(t)
	mockPublisherGateway := mocks.NewMockPublisherGateway(t)
	// mockOrderGateway := mocks.NewMockOrderGateway(t)
	orderId := "1"
	queueName := "my-queue"

	order := &entity.Order{
		ID:          orderId,
		Amount:      10,
		OrderStatus: valueobject.OrderStatus("READY_TO_TAKEOUT"),
	}

	mockDatasource.On("FindOne", "_id", orderId).Return(order, nil)
	mockPublisherGateway.On("GetQueueUrl").Return(queueName)
	mockPublisherGateway.On("PublishMessage", mock.Anything, mock.Anything).Return(nil)

	og := &gateway.OrderGateway{
		Datasource:       mockDatasource,
		OrderApi:         mockOrderApi,
		PublisherGateway: mockPublisherGateway,
	}

	err := og.ReleaseOrder(orderId)

	assert.NoError(t, err)
	mockDatasource.AssertExpectations(t)
	mockPublisherGateway.AssertExpectations(t)
}
func TestOrderGateway_FinishOrder(t *testing.T) {
	mockDatasource := mocks.NewMockDatabaseSource(t)
	mockOrderApi := mocks.NewMockOrderApi(t)
	mockPublisherGateway := mocks.NewMockPublisherGateway(t)
	queueName := "my-queue"
	orderId := "1"

	order := &entity.Order{
		ID:          orderId,
		Amount:      10,
		OrderStatus: valueobject.OrderStatus("COMPLETED"),
	}
	mockDatasource.On("FindOne", "_id", orderId).Return(order, nil)
	mockPublisherGateway.On("GetQueueUrl").Return(queueName)
	mockPublisherGateway.On("PublishMessage", mock.Anything, mock.Anything).Return(nil)

	og := &gateway.OrderGateway{
		Datasource:       mockDatasource,
		OrderApi:         mockOrderApi,
		PublisherGateway: mockPublisherGateway,
	}

	err := og.FinishOrder(orderId)

	assert.NoError(t, err)
	mockDatasource.AssertExpectations(t)
	mockPublisherGateway.AssertExpectations(t)
}
