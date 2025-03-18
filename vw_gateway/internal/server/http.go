package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/redis/go-redis/v9"
	"time"
	captv1 "vw_gateway/api/v1/captcha"
	favorv1 "vw_gateway/api/v1/favorites"
	followv1 "vw_gateway/api/v1/follow"
	idv1 "vw_gateway/api/v1/identity"
	filev1 "vw_gateway/api/v1/userfile"
	infov1 "vw_gateway/api/v1/userinfo"
	"vw_gateway/internal/conf"
	"vw_gateway/internal/pkg/codecs"
	"vw_gateway/internal/pkg/middlewares/auth"
	"vw_gateway/internal/service"
)

func NewWhitelistMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})

	/*以下摘自kratos文档：
	注意：定制中间件是通过 operation 匹配，并不是 http 本身的路由
	operation 是 HTTP 及 gRPC 统一的 gRPC path。
	gRPC path 的拼接规则为 /包名.服务名/方法名(/package.Service/Method)。
	*/
	whiteList["/gateway.api.v1.id.Identity/Register"] = struct{}{}
	whiteList["/gateway.api.v1.id.Identity/Login"] = struct{}{}
	whiteList["/gateway.api.v1.id.Identity/Logout"] = struct{}{}
	whiteList["/gateway.api.v1.captcha.Captcha/GetImageCaptcha"] = struct{}{}
	whiteList["/gateway.api.v1.captcha.Captcha/GetCodeCaptcha"] = struct{}{}
	whiteList["/gateway.api.v1.file.FileService/UploadAvatar"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	jwt *conf.JWT,
	captcha *service.CaptchaService,
	file *service.UserFileService,
	identity *service.UserIdentityService,
	follow *service.FollowService,
	redis *redis.ClusterClient,
	info *service.UserinfoService,
	favorites *service.FavoritesService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			selector.Server(
				auth.JwtAuth(jwt.Secret, time.Duration(jwt.AccessTokenExpiration)*time.Hour, redis),
			).Match(NewWhitelistMatcher()).Build(),
		),
		// 跨域
		http.Filter(handlers.CORS(
			//据AllowedHeaders方法介绍，如果accepting Content-Type除了
			//application/x-www-form-urlencoded, multipart/form-data, 或 text/plain之外还有其他类型，
			//则需要在AllowedHeaders中显式声明“Content-Type”。
			//由于Content-Type可能还会有application/octet-stream等类型，所以这里需要显式声明。
			handlers.AllowedHeaders([]string{"Authorization", "Refresh-Token", "Content-Type"}), // 增加刷新token以及Content-Type的跨域支持.
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "PATCH", "OPTIONS", "DELETE"}),
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
	idv1.RegisterIdentityHTTPServer(srv, identity)
	captv1.RegisterCaptchaHTTPServer(srv, captcha)
	filev1.RegisterFileServiceHTTPServer(srv, file)
	infov1.RegisterUserinfoHTTPServer(srv, info)
	favorv1.RegisterFavoriteHTTPServer(srv, favorites)
	followv1.RegisterFollowHTTPServer(srv, follow)
	return srv
}
