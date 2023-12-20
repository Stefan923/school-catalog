package api

import (
	"backend/api/router"
	"backend/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(config *config.Config) {
	gin.SetMode(config.Server.RunningMode)
	engine := gin.New()

	RegisterRoutes(engine, config)

	err := engine.Run(fmt.Sprintf(":%s", config.Server.InternalPort))
	if err != nil {
		fmt.Println("Error while starting server...")
	}
}

func RegisterRoutes(engine *gin.Engine, config *config.Config) {
	apiRoute := engine.Group("/api")

	v1Route := apiRoute.Group("/v1")
	{
		userAccountsRoute := v1Route.Group("/auth")

		router.StartAuthRouter(userAccountsRoute, config)
	}
}
