package config

import (
	"os"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/constants"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/usecase"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/api/v1/processors"
	mongodb "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/datasource"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/operation/gateway"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitProcessors(dbClient mongo.Client) {

	processorConfig := processors.OrderProcessorConfig{
		Endpoint:        GetQueueProcessorsCfg().OrderEventsQueueEndpoint,
		Region:          GetQueueProcessorsCfg().OrderEventsQueueRegion,
		QueueName:       GetQueueProcessorsCfg().OrderEventsQueue,
		WaitTimeSeconds: 5,
	}

	orderDbAdapter := mongodb.NewMongoAdapter[entity.Order](
		dbClient,
		GetMongoCfg().Database,
		constants.OrderCollection,
	)
	publisherGateway := gateway.NewPublisherGateway(gateway.PublisherGatewayConfig{
		SQSQueueUrl:        GetQueueProcessorsCfg().KitchenEventsQueue,
		SQSEndpoint:        GetQueueProcessorsCfg().KitchenEventsQueueEndpoint,
		AWSRegion:          GetQueueProcessorsCfg().KitchenEventsQueueRegion,
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	})

	orderApi := gateway.NewOrderApi(gateway.OrderApiConfig{OrderApiBaseUrl: GetApiCfg().OrderApiBaseURL})
	gateway := gateway.NewOrderGateway(orderDbAdapter, orderApi, publisherGateway)
	orderUseCase := usecase.NewOrderUseCase(gateway)
	processors.OrderProcessor(processorConfig, orderUseCase)

}
