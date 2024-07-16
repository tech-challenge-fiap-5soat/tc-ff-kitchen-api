package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/usecase"
	valueobject "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/test/mocks"
)

var orderGatewayMock *mocks.MockOrderGateway

func TestOrderUseCase(t *testing.T) {
	t.Parallel()

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

func TestUpdateOrder(t *testing.T) {
	t.Run("should update order successfully", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)

		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.AWAITING_PREPARATION,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		orderToUpdate := entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)
		orderGatewayMock.On("Update", mock.AnythingOfType("*entity.Order")).Return(nil)

		err := orderUseCase.UpdateOrder(orderID, orderToUpdate)

		assert.NoError(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNumberOfCalls(t, "Update", 1)
	})

	t.Run("should return error when order is not found", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		orderToUpdate := entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}

		orderGatewayMock.On("FindById", orderID).Return(nil, errors.New("order not found"))

		err := orderUseCase.UpdateOrder(orderID, orderToUpdate)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNotCalled(t, "Update")
	})

	t.Run("should return error when order status cannot be updated", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.READY_TO_TAKEOUT,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		orderToUpdate := entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)

		err := orderUseCase.UpdateOrder(orderID, orderToUpdate)

		assert.Error(t, err)
		assert.EqualError(t, err, "order cannot be updated cause status is READY_TO_TAKEOUT")
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNotCalled(t, "Update")
	})

	t.Run("should return error when order gateway update fails", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.AWAITING_PREPARATION,
			OrderItems:  []entity.OrderItem{},
			CreatedAt:   valueobject.CustomTime{},
			UpdatedAt:   valueobject.CustomTime{},
			Amount:      0,
		}
		orderToUpdate := entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			CreatedAt:   valueobject.CustomTime{},
			UpdatedAt:   valueobject.CustomTime{},
			Amount:      0,
		}

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)
		orderGatewayMock.On("Update", existentOrder).Return(errors.New("update failed"))

		err := orderUseCase.UpdateOrder(orderID, orderToUpdate)

		assert.Error(t, err)
		assert.EqualError(t, err, "update failed")
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
	})
}
func TestUpdateOrderStatus(t *testing.T) {
	t.Run("should update order status successfully", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)

		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		newStatus := valueobject.READY_TO_TAKEOUT

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)
		orderGatewayMock.On("Update", mock.AnythingOfType("*entity.Order")).Return(nil)
		orderGatewayMock.On("ReleaseOrder", orderID).Return(nil)

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.NoError(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertCalled(t, "Update", mock.AnythingOfType("*entity.Order"))
		orderGatewayMock.AssertCalled(t, "ReleaseOrder", orderID)
	})

	t.Run("should return error when order is not found", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		newStatus := valueobject.READY_TO_TAKEOUT

		orderGatewayMock.On("FindById", orderID).Return(nil, errors.New("order not found"))

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNotCalled(t, "Update", mock.AnythingOfType("*entity.Order"))
		orderGatewayMock.AssertNotCalled(t, "ReleaseOrder", orderID)
	})

	t.Run("should return error when updating to previous status", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		newStatus := valueobject.AWAITING_PREPARATION

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNotCalled(t, "Update", mock.AnythingOfType("*entity.Order"))
		orderGatewayMock.AssertNotCalled(t, "ReleaseOrder", orderID)
	})

	t.Run("should return error when updating to invalid next status", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		newStatus := valueobject.COMPLETED

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNotCalled(t, "Update", mock.AnythingOfType("*entity.Order"))
		orderGatewayMock.AssertNotCalled(t, "ReleaseOrder", orderID)
	})

	t.Run("should handle error when updating order status", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		newStatus := valueobject.READY_TO_TAKEOUT

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)
		orderGatewayMock.On("Update", mock.AnythingOfType("*entity.Order")).Return(errors.New("update error"))

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertCalled(t, "Update", mock.AnythingOfType("*entity.Order"))
		orderGatewayMock.AssertNotCalled(t, "ReleaseOrder", orderID)
	})

	t.Run("should handle error when releasing order", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)

		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.PREPARING,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		newStatus := valueobject.READY_TO_TAKEOUT

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)
		orderGatewayMock.On("Update", mock.AnythingOfType("*entity.Order")).Return(nil)
		orderGatewayMock.On("ReleaseOrder", orderID).Return(errors.New("release error"))

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertCalled(t, "Update", mock.AnythingOfType("*entity.Order"))
		orderGatewayMock.AssertCalled(t, "ReleaseOrder", orderID)
	})

	t.Run("should update order status successfully", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)

		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.AWAITING_PREPARATION,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		newStatus := valueobject.PREPARING

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)
		orderGatewayMock.On("Update", mock.AnythingOfType("*entity.Order")).Return(nil)

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.NoError(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNumberOfCalls(t, "Update", 1)
	})

	t.Run("should return error when order is not found", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"

		orderGatewayMock.On("FindById", orderID).Return(nil, errors.New("order not found"))

		err := orderUseCase.UpdateOrderStatus(orderID, valueobject.PREPARING)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNotCalled(t, "Update")
	})

	t.Run("should return error when order status cannot be updated", func(t *testing.T) {
		orderGatewayMock := mocks.NewMockOrderGateway(t)
		orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)
		orderID := "123"
		existentOrder := &entity.Order{
			ID:          orderID,
			OrderStatus: valueobject.COMPLETED,
			OrderItems:  []entity.OrderItem{},
			Amount:      0,
		}
		newStatus := valueobject.PREPARING

		orderGatewayMock.On("FindById", orderID).Return(existentOrder, nil)

		err := orderUseCase.UpdateOrderStatus(orderID, newStatus)

		assert.Error(t, err)
		orderGatewayMock.AssertCalled(t, "FindById", orderID)
		orderGatewayMock.AssertNotCalled(t, "Update")
	})

}

func TestCreateOrder(t *testing.T) {
	orderGatewayMock := mocks.NewMockOrderGateway(t)
	orderUseCase := usecase.NewOrderUseCase(orderGatewayMock)

	order := entity.Order{
		ID:          "123",
		OrderStatus: valueobject.AWAITING_PREPARATION,
		OrderItems:  []entity.OrderItem{},
		Amount:      0,
		Customer:    entity.Customer{},
	}

	orderGatewayMock.On("Save", mock.AnythingOfType("*entity.Order")).Return("123", nil)

	orderID, err := orderUseCase.CreateOrder(order)

	assert.NoError(t, err)
	assert.Equal(t, "123", orderID)
	orderGatewayMock.AssertCalled(t, "Save", &order)
}
