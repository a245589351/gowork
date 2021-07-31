package initialize

import (
	"github.com/gin-gonic/gin"

	"gowork/mission2/router"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()

	ApiGroup := Router.Group("/h/v1")
	router.InitHelloRouter(ApiGroup)

	return Router
}
