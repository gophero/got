package redist

import (
	"context"
	"fmt"
	"log"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	redismock "github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
)

func MockRedis() (*redis.Client, redismock.ClientMock) {
	return redismock.NewClientMock()
}

func NewMiniRedis() *redis.Client {
	// 测试用miniredis
	mr, err := miniredis.Run()
	if err != nil {
		panic(fmt.Errorf("new test redis error: %v", err))
	}
	// 使用miniredis创建client
	client := redis.NewClient(&redis.Options{
		Addr:         mr.Addr(),
		Password:     "",
		DB:           0,
		MaxRetries:   5,
		MinIdleConns: 2,
		TLSConfig:    nil,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis error: %s", err.Error()))
	}
	log.Printf("redis connected, url: %s\n", client.Conn().String())
	return client
}

func NewRedisCluster() redis.UniversalClient {
	// TODO mock redis cluster
	return nil
}
