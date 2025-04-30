package videodata

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"vw_gateway/internal/conf"
	userinfov1 "vw_user/api/v1/userinfo"
	videointeractv1 "vw_video/api/v1/interact"

	"github.com/go-kratos/kratos/v2/log"
	videoinfov1 "vw_video/api/v1/videoinfo"
)

const (
	kb = 1024
	mb = 1024 * kb

	videoService = "videoweb.video.service"
	userService  = "videoweb.user.service"
)

// Data .
type Data struct {
	log                 *log.Helper
	videoInfoClient     videoinfov1.VideoInfoClient
	videoInteractClient videointeractv1.VideoInteractClient
	userInfoClient      userinfov1.UserinfoClient
	redis               *redis.ClusterClient
	dtmServerAddr       string
}

// NewData .
func NewData(
	logger log.Logger,
	videoInfoClient videoinfov1.VideoInfoClient,
	userInfoClient userinfov1.UserinfoClient,
	videoInteractClient videointeractv1.VideoInteractClient,
	redisCluster *redis.ClusterClient,
	dtm *conf.DTM,

) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		log:                 log.NewHelper(logger),
		videoInfoClient:     videoInfoClient,
		videoInteractClient: videoInteractClient,
		userInfoClient:      userInfoClient,
		redis:               redisCluster,
		dtmServerAddr:       "discovery:///dtm",
	}, cleanup, nil
}

func NewVideoInfoClient(r registry.Discovery, s *conf.Service) videoinfov1.VideoInfoClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(s.Video.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return videoinfov1.NewVideoInfoClient(conn)
}

func NewVideoInteractClient(r registry.Discovery, s *conf.Service) videointeractv1.VideoInteractClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(s.Video.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return videointeractv1.NewVideoInteractClient(conn)
}

func NewUserInfoClient(r registry.Discovery, s *conf.Service) userinfov1.UserinfoClient {
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
	return userinfov1.NewUserinfoClient(conn)
}
