package core

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/sync/errgroup"

	"gowork/mission2/initialize"
	"gowork/mission2/proto/hello"
	"gowork/mission2/server"
	"gowork/mission2/server/grpc"
	"gowork/mission2/server/http"
	"gowork/mission2/service"
)

func Run() error {
	// 启动客户端与服务端
	eg, ctx := errgroup.WithContext(context.Background())
	wg := sync.WaitGroup{}
	var servers []server.Server
	httpServer := http.NewServer(http.Address(":8000"), http.Router(initialize.InitRouters()))
	grpcServer := grpc.NewServer(grpc.Address(":9000"))
	hello.RegisterHelloServer(grpcServer.Server, &service.Hello{})
	servers = append(servers, httpServer)
	servers = append(servers, grpcServer)

	for _, srv := range servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done()
			return srv.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}

	// 连接服务端
	eg.Go(func() error {
		wg.Add(1)
		err := initialize.InitSrvConn()
		wg.Done()
		return err
	})

	wg.Wait()

	// 监听系统信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-quit:
				return errors.New("quit")
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}
