package core

import (
	"context"
	"errors"
	"gowork/week4/global"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"

	"gowork/week4/initialize"
)

func Run() error {
	// 启动客户端与服务端
	eg, ctx := errgroup.WithContext(context.Background())
	wg := sync.WaitGroup{}

	for _, srv := range global.Servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done()
			stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return srv.Stop(stopCtx)
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
