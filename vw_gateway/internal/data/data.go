package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"vw_gateway/internal/conf"
	captv1 "vw_user/api/v1/captcha"
	idv1 "vw_user/api/v1/identity"
	filev1 "vw_user/api/v1/userfile"
	infov1 "vw_user/api/v1/userinfo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistrar,
	NewDiscovery,
	NewUserIdentityRepo,
	NewCaptchaRepo,
	NewUserFileRepo,
	NewUserIdentityClient,
	NewUserinfoClient,
	NewCaptchaClient,
	NewRedisClusterClient,
	NewFileClient,
	NewUserInfoRepo,
)

// Data .
type Data struct {
	log                *log.Helper
	userIdentityClient idv1.IdentityClient
	userInfoClient     infov1.UserinfoClient
	captchaClient      captv1.CaptchaClient
	fileClient         filev1.FileServiceClient
	redis              *redis.ClusterClient
}

// NewData .
func NewData(
	logger log.Logger,
	identityClient idv1.IdentityClient,
	infoClient infov1.UserinfoClient,
	captchaClient captv1.CaptchaClient,
	fileClient filev1.FileServiceClient,
	redisCluster *redis.ClusterClient,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		log:                log.NewHelper(logger),
		userIdentityClient: identityClient,
		userInfoClient:     infoClient,
		captchaClient:      captchaClient,
		fileClient:         fileClient,
		redis:              redisCluster,
	}, cleanup, nil
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

func NewUserIdentityClient(r registry.Discovery, s *conf.Service) idv1.IdentityClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(s.User.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return idv1.NewIdentityClient(conn)
}

func NewUserinfoClient(r registry.Discovery, s *conf.Service) infov1.UserinfoClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(s.User.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return infov1.NewUserinfoClient(conn)
}

func NewCaptchaClient(r registry.Discovery, s *conf.Service) captv1.CaptchaClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(s.User.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return captv1.NewCaptchaClient(conn)
}

func NewFileClient(r registry.Discovery, s *conf.Service) filev1.FileServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(s.User.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return filev1.NewFileServiceClient(conn)
}
