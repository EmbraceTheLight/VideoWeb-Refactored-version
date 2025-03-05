package data_test

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"vw_gateway/internal/conf"
	"vw_user/api/v1/captcha"
)

func TestClient(t *testing.T) {
	c := consulAPI.DefaultConfig()
	c.Address = "127.0.0.1:8500"
	c.Scheme = "http"
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///videoweb.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	require.NoError(t, err)
	require.NotNil(t, conn)

	client := captv1.NewCaptchaClient(conn)
	_, err = client.DeleteCodeFromCache(context.Background(), &captv1.DeleteCodeFromCacheReq{
		Email: "1010642166@qq.com",
	})
	require.NoError(t, err)
}

func TestNewRedisCluster(t *testing.T) {
	c := config.New(
		config.WithSource(
			file.NewSource("../../configs/config.yaml"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		require.NoError(t, err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		require.NoError(t, err)
	}

	var address []string
	ipAddress := bc.Data.RedisCluster.Host
	for _, port := range bc.Data.RedisCluster.Port {
		address = append(address, fmt.Sprintf("%s:%s", ipAddress, port))
	}

	redisCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        address,
		Password:     bc.Data.RedisCluster.Password,
		PoolSize:     int(bc.Data.RedisCluster.PoolSize),
		MinIdleConns: int(bc.Data.RedisCluster.MinIdleConns),
		MaxRetries:   int(bc.Data.RedisCluster.MaxRetries),
		DialTimeout:  bc.Data.RedisCluster.DialTimeout.AsDuration(),
		ReadTimeout:  bc.Data.RedisCluster.ReadTimeout.AsDuration(),
		WriteTimeout: bc.Data.RedisCluster.WriteTimeout.AsDuration(),
		PoolTimeout:  bc.Data.RedisCluster.PoolTimeout.AsDuration(),
	})
	err := redisCluster.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("redis cluster test...")
	fmt.Println("redis cluster Set key-values")
	redisCluster.Set(ctx, "111", "12121", 30*time.Minute)
	redisCluster.Set(ctx, "hello", "world", 30*time.Minute)

	fmt.Println("redis cluster get key-values")
	res1, err := redisCluster.Get(ctx, "111").Result()
	require.NoError(t, err)
	fmt.Println("value of key `111` is: ", res1)

	res2, err := redisCluster.Get(ctx, "hello").Result()
	fmt.Println("value of key `hello` is: ", res2)

	require.NoError(t, err)

}
