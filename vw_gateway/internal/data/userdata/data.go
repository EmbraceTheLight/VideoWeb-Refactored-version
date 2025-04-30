package userdata

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"vw_gateway/internal/conf"
	captv1 "vw_user/api/v1/captcha"
	favorv1 "vw_user/api/v1/favorites"
	followv1 "vw_user/api/v1/follow"
	idv1 "vw_user/api/v1/identity"
	filev1 "vw_user/api/v1/userfile"
	infov1 "vw_user/api/v1/userinfo"

	"github.com/go-kratos/kratos/v2/log"
)

// Data .
type Data struct {
	log                *log.Helper
	userIdentityClient idv1.IdentityClient
	userInfoClient     infov1.UserinfoClient
	captchaClient      captv1.CaptchaClient
	fileClient         filev1.FileServiceClient
	favoritesClient    favorv1.FavoriteClient
	followClient       followv1.FollowClient

	redis         *redis.ClusterClient
	dtmServerAddr string
}

// NewData .
func NewData(
	logger log.Logger,
	identityClient idv1.IdentityClient,
	infoClient infov1.UserinfoClient,
	captchaClient captv1.CaptchaClient,
	fileClient filev1.FileServiceClient,
	favoritesClient favorv1.FavoriteClient,
	followClient followv1.FollowClient,
	redisCluster *redis.ClusterClient,
	dtm *conf.DTM,
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
		favoritesClient:    favoritesClient,
		followClient:       followClient,
		redis:              redisCluster,
		dtmServerAddr:      dtm.DtmConfig.Target,
	}, cleanup, nil
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

func NewFavoritesClient(r registry.Discovery, s *conf.Service) favorv1.FavoriteClient {
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
	return favorv1.NewFavoriteClient(conn)
}

func NewFollowClient(r registry.Discovery, s *conf.Service) followv1.FollowClient {
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
	return followv1.NewFollowClient(conn)
}
