package router

import (
	"github.com/gin-gonic/gin"

	"gowork/mission2/handler"
)

func InitHelloRouter(Router *gin.RouterGroup) {
	HelloRouter := Router.Group("hello")
	{
		HelloRouter.GET("", handler.Hello)
	}
}
