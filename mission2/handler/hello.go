package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"gowork/mission2/global"
	"gowork/mission2/proto/hello"
)

func Hello(ctx *gin.Context) {
	r, err := global.HelloSrvClient.SayHello(context.Background(), &hello.SayHelloRequest{Message: "VE!!"})

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, r)
}
