package router

import (
	"gowork/week4/middleware"
	"time"

	"github.com/gin-gonic/gin"

	"gowork/week4/handler"
)

func InitHelloRouter(Router *gin.RouterGroup) {
	HelloRouter := Router.Group("hello").Use(middleware.TimeoutMiddleware(time.Second * 2))
	{
		HelloRouter.GET("", handler.Hello)
	}
}
