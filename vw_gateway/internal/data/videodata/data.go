package videodata

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"vw_gateway/internal/conf"
	videoinfov1 "vw_video/api/v1/videoinfo"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	kb = 1024
	mb = 1024 * kb
)

// Data .
type Data struct {
	log             *log.Helper
	videoInfoClient videoinfov1.VideoInfoClient
	redis           *redis.ClusterClient
}

// NewData .
func NewData(
	logger log.Logger,
	videoInfoClient videoinfov1.VideoInfoClient,
	redisCluster *redis.ClusterClient,
) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		log:             log.NewHelper(logger),
		videoInfoClient: videoInfoClient,
		redis:           redisCluster,
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
