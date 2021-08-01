package router

import (
	"gowork/mission2/middleware"
	"time"

	"github.com/gin-gonic/gin"

	"gowork/mission2/handler"
)

func InitHelloRouter(Router *gin.RouterGroup) {
	HelloRouter := Router.Group("hello").Use(middleware.TimeoutMiddleware(time.Second * 2))
	{
		HelloRouter.GET("", handler.Hello)
	}
}
