package processors

import (
	"fmt"

	"github.com/inaciogu/go-sqs/consumer"
	"github.com/inaciogu/go-sqs/consumer/message"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
)

type OrderProcessorConfig struct {
	Endpoint        string
	QueueName       string
	Region          string
	WaitTimeSeconds int
}

type OrderHandler struct {
	OrderUseCase interfaces.OrderUseCase
}

func OrderProcessor(config OrderProcessorConfig, orderUseCase interfaces.OrderUseCase) {

	handler := OrderHandler{
		OrderUseCase: orderUseCase,
	}

	clientOptions := consumer.SQSClientOptions{
		Region:          config.Region,
		QueueName:       config.QueueName,
		Handle:          handler.orderEventHandler,
		WaitTimeSeconds: int64(config.WaitTimeSeconds),
	}
	if config.Endpoint != "" {
		clientOptions.Endpoint = config.Endpoint
	}
	go consumer.New(nil, clientOptions).Start()
}

func (op OrderHandler) orderEventHandler(message *message.Message) bool {
	orderEvent := entity.OrderEvent{}

	err := message.Unmarshal(&orderEvent)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Print("Order Event Received: ", orderEvent)
	fmt.Print("Order Event: ", orderEvent.EventType)

	if orderEvent.EventType == "RequestOrderPreparation" {
		op.OrderUseCase.CreateOrder(*orderEvent.Order)
		return true
	}
	return false
}
