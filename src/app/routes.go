package app

import (
	"github.com/AdieOlami/auth-service/src/controller"
	"github.com/AdieOlami/auth-service/src/data"
	"github.com/AdieOlami/auth-service/src/services"
)

func (a *App) initializeRoutes() {
	atHandler := controller.NewController(services.NewService(a.Cassandra, data.NewTokenRepository(), data.NewUsersRepository()))
	v1 := a.Router.Group("/api/v1")
	{
		//Auth routes
		v1.GET("/auth/access_token/:access_token_id", atHandler.AuthenticateRequest)
		v1.POST("/auth/access_token", atHandler.Create)
	}
}
