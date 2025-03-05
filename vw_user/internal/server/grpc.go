package server

import (
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	captv1 "vw_user/api/v1/captcha"
	idv1 "vw_user/api/v1/identity"
	filev1 "vw_user/api/v1/userfile"
	infov1 "vw_user/api/v1/userinfo"
	"vw_user/internal/conf"
	"vw_user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	identity *service.UserIdentityService,
	info *service.UserInfoService,
	captcha *service.CaptchaService,
	file *service.FileService,
	logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	captv1.RegisterCaptchaServer(srv, captcha)
	idv1.RegisterIdentityServer(srv, identity)
	infov1.RegisterUserinfoServer(srv, info)
	filev1.RegisterFileServiceServer(srv, file)
	return srv
}
