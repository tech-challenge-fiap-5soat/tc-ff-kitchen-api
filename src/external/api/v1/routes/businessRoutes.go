package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/constants"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/core/entity"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/api/infra/config"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/api/v1/handlers"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/operation/controller"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/operation/gateway"

	mongodb "github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/datasource"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterBusinessRoutes(gServer *gin.RouterGroup, dbClient mongo.Client) {
	groupServer := gServer.Group("/v1")

	registerOrderHandler(groupServer, dbClient)
}

func registerOrderHandler(groupServer *gin.RouterGroup, dbClient mongo.Client) {
	orderDbAdapter := mongodb.NewMongoAdapter[entity.Order](
		dbClient,
		config.GetMongoCfg().Database,
		constants.OrderCollection,
	)

	orderApi := gateway.NewOrderApi(gateway.OrderApiConfig{OrderApiBaseUrl: config.GetApiCfg().OrderApiBaseURL})

	orderInteractor := controller.NewOrderController(orderDbAdapter, orderApi)
	handlers.NewOrderHandler(groupServer, orderInteractor)
}
