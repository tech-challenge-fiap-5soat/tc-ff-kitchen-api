package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/api/infra/config"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/api/v1/middlewares"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/external/api/v1/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(gServer *gin.Engine, dbClient mongo.Client) {
	gServer.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/liveness", "/health/readiness"),
		middlewares.CORSMiddleware(),
		middlewares.CheckAccessToken("/api/v1/customer/authorization", "/health/liveness", "/health/readiness", "/docs/*"),
		gin.Recovery(),
	)

	RegisterHealthRoutes(gServer)
	RegisterSwaggerRoutes(gServer)

	api := gServer.Group("/api")
	routes.RegisterBusinessRoutes(api, dbClient)

	gServer.Run(fmt.Sprintf(":%s", config.GetApiCfg().Port))
}
