package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

var (
	rdb *redis.Client
)

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   7,
	})
	rdb.AddHook(redisotel.TracingHook{})
	ctx := context.Background()

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	batchCreate(ctx, 5120)
}

func batchCreate(ctx context.Context, i int64) {
	var (
		ds []string
		sb []byte
		k  int64
	)

	for k = 0; k < i; k++ {
		sb = append(sb, 'a')
	}

	for j := 0; j < 10000; j++ {
		ds = append(ds, cast.ToString(j))
		ds = append(ds, string(sb))
	}

	rdb.MSet(ctx, ds)

	fmt.Println("插入完毕")
}
