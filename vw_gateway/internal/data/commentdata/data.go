package commentdata

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	vcommentv1 "vw_comment/api/v1/video_comment"
	"vw_gateway/internal/conf"
)

type Data struct {
	videoCommentClient vcommentv1.VideoCommentClient

	redis *redis.ClusterClient
	log   *log.Helper
}

func NewData(videoCommentClient vcommentv1.VideoCommentClient, redis *redis.ClusterClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		videoCommentClient: videoCommentClient,
		redis:              redis,
		log:                log.NewHelper(logger),
	}, cleanup, nil
}

func NewVideoCommentClient(r registry.Discovery, s *conf.Service) vcommentv1.VideoCommentClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(s.Comment.Endpoint),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return vcommentv1.NewVideoCommentClient(conn)
}
