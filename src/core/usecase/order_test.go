package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/usecase"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"

	"github.com/stretchr/testify/assert"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/test/mocks"
)

var orderGatewayMock *mocks.MockOrderGateway

func TestOrderUseCase(t *testing.T) {
	// t.Parallel()

	t.Run("should return order by given id", func(t *testing.T) {
		expectedOrder := &entity.Order{ID: "123"}

		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindById", "123").Return(expectedOrder, nil)

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrder, err := useCase.FindById("123")

		assert.Nil(t, err)
		assert.NotNil(t, resultOrder)
	})

	t.Run("should return empty result when not found order by id", func(t *testing.T) {
		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindById", "123").Return(nil, nil)

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrder, err := useCase.FindById("123")

		assert.NoError(t, err)
		assert.Nil(t, resultOrder)
	})

	t.Run("should return error in Repository when call FindById", func(t *testing.T) {
		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindById", "789").Return(nil, errors.New("repository error"))

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		result, err := useCase.FindById("789")

		assert.Error(t, err)
		assert.Nil(t, result)
		orderGatewayMock.AssertExpectations(t)
	})

	t.Run("should return orders by status", func(t *testing.T) {
		expectedOrders := []entity.Order{
			{ID: "1", OrderStatus: valueobject.AWAITING_PREPARATION},
		}

		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindAllByStatus", valueobject.AWAITING_PREPARATION).Return(expectedOrders, nil)

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrders, err := useCase.GetAllByStatus(valueobject.AWAITING_PREPARATION)

		assert.NoError(t, err)
		assert.Len(t, resultOrders, len(expectedOrders))
	})

	t.Run("should return empty list when not found orders by status", func(t *testing.T) {
		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindAllByStatus", valueobject.COMPLETED).Return([]entity.Order{}, nil)

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrders, err := useCase.GetAllByStatus(valueobject.COMPLETED)

		assert.NoError(t, err)
		assert.Empty(t, resultOrders)
	})

	t.Run("should handle repository error", func(t *testing.T) {
		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindAllByStatus", valueobject.READY_TO_TAKEOUT).Return(nil, errors.New("repository error"))

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrders, err := useCase.GetAllByStatus(valueobject.READY_TO_TAKEOUT)

		assert.Error(t, err)
		assert.Nil(t, resultOrders)
	})

	t.Run("should return all orders sorted by READY_TO_TAKEOUT > PREPARING", func(t *testing.T) {
		expectedOrders := []entity.Order{
			{ID: "1", OrderStatus: valueobject.PREPARING},
			{ID: "3", OrderStatus: valueobject.READY_TO_TAKEOUT},
			{ID: "4", OrderStatus: valueobject.READY_TO_TAKEOUT},
			{ID: "6", OrderStatus: valueobject.PREPARING},
		}

		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindAll").Return(expectedOrders, nil)

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrders, err := useCase.FindAll()

		assert.NoError(t, err)
		assert.Len(t, resultOrders, len(expectedOrders))
		assert.Equal(t, valueobject.READY_TO_TAKEOUT, resultOrders[0].OrderStatus)
		assert.Equal(t, valueobject.READY_TO_TAKEOUT, resultOrders[1].OrderStatus)
		assert.Equal(t, valueobject.PREPARING, resultOrders[2].OrderStatus)
		assert.Equal(t, valueobject.PREPARING, resultOrders[3].OrderStatus)
	})

	t.Run("should return all orders sorted by createdAt", func(t *testing.T) {
		currentTime := time.Now()

		expectedOrders := []entity.Order{
			{ID: "1", OrderStatus: valueobject.READY_TO_TAKEOUT, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(2) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
			{ID: "2", OrderStatus: valueobject.READY_TO_TAKEOUT, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(1) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
			{ID: "3", OrderStatus: valueobject.PREPARING, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(4) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
			{ID: "4", OrderStatus: valueobject.PREPARING, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(3) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
		}

		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindAll").Return(expectedOrders, nil)

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrders, err := useCase.FindAll()

		assert.NoError(t, err)
		assert.Len(t, resultOrders, len(expectedOrders))
		assert.Equal(t, resultOrders[0].OrderStatus, valueobject.READY_TO_TAKEOUT)
		assert.Equal(t, resultOrders[1].OrderStatus, valueobject.READY_TO_TAKEOUT)
		assert.Equal(t, resultOrders[2].OrderStatus, valueobject.PREPARING)
		assert.Equal(t, resultOrders[3].OrderStatus, valueobject.PREPARING)
		assert.True(t, resultOrders[0].CreatedAt.Before(resultOrders[1].CreatedAt.Time))
		assert.True(t, resultOrders[1].CreatedAt.Before(resultOrders[2].CreatedAt.Time))
		assert.True(t, resultOrders[2].CreatedAt.Before(resultOrders[3].CreatedAt.Time))
	})

	t.Run("should return all orders without COMPLETED status", func(t *testing.T) {
		currentTime := time.Now()

		expectedOrders := []entity.Order{
			{ID: "1", OrderStatus: valueobject.READY_TO_TAKEOUT, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(2) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
			{ID: "2", OrderStatus: valueobject.READY_TO_TAKEOUT, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(1) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
			{ID: "3", OrderStatus: valueobject.PREPARING, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(4) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
			{ID: "4", OrderStatus: valueobject.COMPLETED, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(4) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
			{ID: "5", OrderStatus: valueobject.PREPARING, CreatedAt: valueobject.CustomTime{
				Time: currentTime.Add(
					time.Hour*time.Duration(3) +
						time.Minute*time.Duration(0) +
						time.Second*time.Duration(0),
				),
			}},
		}

		orderGatewayMock = mocks.NewMockOrderGateway(t)
		orderGatewayMock.On("FindAll").Return(expectedOrders, nil)

		useCase := usecase.NewOrderUseCase(orderGatewayMock)

		resultOrders, err := useCase.FindAll()

		assert.NoError(t, err)
		assert.Len(t, resultOrders, len(expectedOrders)-1)

		for _, order := range resultOrders {
			assert.NotEqual(t, valueobject.COMPLETED, order.OrderStatus)
		}

		assert.Equal(t, resultOrders[0].OrderStatus, valueobject.READY_TO_TAKEOUT)
		assert.Equal(t, resultOrders[1].OrderStatus, valueobject.READY_TO_TAKEOUT)
		assert.Equal(t, resultOrders[2].OrderStatus, valueobject.PREPARING)
		assert.Equal(t, resultOrders[3].OrderStatus, valueobject.PREPARING)
		assert.True(t, resultOrders[0].CreatedAt.Before(resultOrders[1].CreatedAt.Time))
		assert.True(t, resultOrders[1].CreatedAt.Before(resultOrders[2].CreatedAt.Time))
		assert.True(t, resultOrders[2].CreatedAt.Before(resultOrders[3].CreatedAt.Time))
	})
}
