package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gowork/mission2/global"
	"gowork/mission2/proto/hello"
)

func Hello(ctx *gin.Context) {
	r, err := global.HelloSrvClient.SayHello(ctx, &hello.SayHelloRequest{Message: "VE!!"})

	if err != nil {
		panic(err)
	}

	fmt.Println("test...http客户端接收到grpc返回")
	ctx.JSON(http.StatusOK, r)
}
