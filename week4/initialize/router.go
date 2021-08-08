package initialize

import (
	"github.com/gin-gonic/gin"

	"gowork/week4/router"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("/h/v1")
	router.InitHelloRouter(ApiGroup)

	return Router
}
