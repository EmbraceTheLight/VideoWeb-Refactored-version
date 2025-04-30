package data

import (
	"context"
	"fmt"
	_ "github.com/dtm-labs/driver-kratos"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"vw_gateway/internal/conf"
	"vw_gateway/internal/data/userdata"
	"vw_gateway/internal/data/videodata"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewRegistrar,
	NewDiscovery,
	NewRedisClusterClient,

	userdata.NewData,
	userdata.NewUserIdentityRepo,
	userdata.NewCaptchaRepo,
	userdata.NewUserFileRepo,
	userdata.NewFollowRepo,
	userdata.NewUserInfoRepo,
	userdata.NewFavoritesRepo,
	userdata.NewUserIdentityClient,
	userdata.NewUserinfoClient,
	userdata.NewCaptchaClient,
	userdata.NewFileClient,
	userdata.NewFavoritesClient,
	userdata.NewFollowClient,

	videodata.NewData,
	videodata.NewVideoInfoRepo,
	videodata.NewInteractRepo,
	videodata.NewVideoInfoClient,
	videodata.NewVideoInteractClient,
)

func NewRedisClusterClient(c *conf.Data) *redis.ClusterClient {
	var address []string
	ipAddress := c.RedisCluster.Host
	for _, port := range c.RedisCluster.Port {
		address = append(address, fmt.Sprintf("%s:%s", ipAddress, port))
	}
	redisCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        address,
		Password:     c.RedisCluster.Password,
		PoolSize:     int(c.RedisCluster.PoolSize),
		MinIdleConns: int(c.RedisCluster.MinIdleConns),
		MaxRetries:   int(c.RedisCluster.MaxRetries),
		DialTimeout:  c.RedisCluster.DialTimeout.AsDuration(),
		ReadTimeout:  c.RedisCluster.ReadTimeout.AsDuration(),
		WriteTimeout: c.RedisCluster.WriteTimeout.AsDuration(),
		PoolTimeout:  c.RedisCluster.PoolTimeout.AsDuration(),
	})
	err := redisCluster.ForEachShard(context.Background(), func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	return redisCluster
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
