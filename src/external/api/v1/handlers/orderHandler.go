package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	orderStatus "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/valueObject"
)

type orderHandler struct {
	interactor interfaces.OrderController
}

const (
	fieldErrMsg = "Invalid value for field: '%s'"
)

func NewOrderHandler(gRouter *gin.RouterGroup, interactor interfaces.OrderController) {
	handler := &orderHandler{
		interactor: interactor,
	}

	gRouter.GET("/order/:id", handler.FindByIdHandler)
	gRouter.POST("/prepare-order", handler.PrepareOrderHandler)
	gRouter.PUT("/order/:id", handler.UpdateOrderHandler)
	gRouter.PUT("/order/:id/status/:status", handler.UpdateStatusOrderHandler)
}

// Get Order godoc
// @Summary Get order by ID
// @Description Get order by ID
// @Tags Order Routes
// @Param        id   path      string  true  "Order ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} entity.Order{}
// @Router /api/v1/order/{id} [get]
func (handler *orderHandler) FindByIdHandler(c *gin.Context) {
	orderId, exists := c.Params.Get("id")
	var order *entity.Order // Only to swaggo doc

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	result, err := handler.interactor.FindById(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order = result

	c.JSON(http.StatusOK, order)
}

// Create Order godoc
// @Summary Prepare new order
// @Description Prepare new order
// @Tags Order Routes
// @Param        data   body      dto.OrderCreateDTO  true  "Order information and customer CPF"
// @Accept  json
// @Produce  json
// @Success 200 {object} interface{}
// @Router /api/v1/order [post]
func (handler *orderHandler) PrepareOrderHandler(c *gin.Context) {
	var order entity.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		var verr validator.ValidationErrors
		var msgFieldError string
		if errors.As(err, &verr) {
			msgFieldError = strings.Split(verr[0].Namespace(), ".")[1]
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(fieldErrMsg, msgFieldError)})
			return
		}
	}

	if len(order.OrderItems) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "at least one product must be selected"})
		return
	}

	orderId, err := handler.interactor.CreateOrder(order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"OrderId": orderId})
}

// Update Order godoc
// @Summary Update order
// @Description Update order
// @Tags Order Routes
// @Param        data   body      dto.OrderUpdateDTO  true  "Order information and customer CPF"
// @Accept  json
// @Produce  json
// @Success 204
// @Router /api/v1/order/{id} [put]
func (handler *orderHandler) UpdateOrderHandler(c *gin.Context) {
	orderId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	var order entity.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		var verr validator.ValidationErrors
		var msgFieldError string
		if errors.As(err, &verr) {
			msgFieldError = strings.Split(verr[0].Namespace(), ".")[1]
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(fieldErrMsg, msgFieldError)})
			return
		}
	}

	if len(order.OrderItems) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "at least one product must be selected"})
		return
	}

	err := handler.interactor.UpdateOrder(orderId, order)

	if err != nil && err.Error() == "record not found" {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// Update Order Status godoc
// @Summary Update order status
// @Description Update order status
// @Tags Order Routes
// @Param        id   path  string  true  "Order ID"
// @Param        status   path  string  true  "AWAITING_PREPARATION, PREPARING, READY_TO_TAKEOUT or COMPLETED"
// @Accept  json
// @Produce  json
// @Success 204
// @Router /api/v1/order/{id}/status/{status} [put]
func (handler orderHandler) UpdateStatusOrderHandler(c *gin.Context) {
	orderId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	status, exists := c.Params.Get("status")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	sts, err := orderStatus.ParseOrderStatus(status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	err = handler.interactor.UpdateOrderStatus(orderId, sts)

	if err != nil && strings.Contains(err.Error(), "record not found") {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	if err != nil && strings.Contains(err.Error(), "cannot be updated to") {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if err != nil && strings.Contains(err.Error(), "previous status") {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
