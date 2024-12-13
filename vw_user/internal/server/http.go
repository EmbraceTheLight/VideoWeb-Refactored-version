package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/gorilla/handlers"
	"time"
	v1id "vw_user/api/user/v1/identity"
	v1info "vw_user/api/user/v1/userinfo"
	"vw_user/internal/conf"
	"vw_user/internal/pkg/codecs"
	"vw_user/internal/pkg/middlewares/auth"
	"vw_user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewWhitelistMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})

	/*以下摘自kratos文档：
	注意：定制中间件是通过 operation 匹配，并不是 http 本身的路由
	operation 是 HTTP 及 gRPC 统一的 gRPC path。
	gRPC path 的拼接规则为 /包名.服务名/方法名(/package.Service/Method)。
	*/
	whiteList["/user.v1.id.Identity/Register"] = struct{}{}
	whiteList["/user.v1.id.Identity/Login"] = struct{}{}
	whiteList["/user.v1.id.Captcha/GetImageCaptcha"] = struct{}{}
	whiteList["/user.v1.id.Captcha/GetCodeCaptcha"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	identity *service.UserIdentityService,
	info *service.UserInfoService,
	jwt *conf.JWT,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		//中间件顺序：先进后出
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			// jwt鉴权中间件
			selector.Server(
				auth.JwtAuth(jwt.Secret, time.Duration(jwt.AccessTokenExpiration)*time.Hour),
			).Match(NewWhitelistMatcher()).Build(),
		),
		// 跨域
		http.Filter(handlers.CORS(
			//据AllowedHeaders方法介绍，如果accepting Content-Type除了
			//application/x-www-form-urlencoded, multipart/form-data, 或 text/plain之外还有其他类型，
			//则需要在AllowedHeaders中显式声明“Content-Type”。
			//由于Content-Type可能还会有application/octet-stream等类型，所以这里需要显式声明。
			handlers.AllowedHeaders([]string{"Authorization", "Refresh-Token", "Content-Type"}), // 增加刷新token以及Content-Type的跨域支持.
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
		http.RequestDecoder(codecs.RequestDecoder),
		http.ErrorEncoder(codecs.ErrorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1id.RegisterIdentityHTTPServer(srv, identity)
	v1id.RegisterCaptchaHTTPServer(srv, identity)
	v1info.RegisterUserinfoHTTPServer(srv, info)
	return srv
}
