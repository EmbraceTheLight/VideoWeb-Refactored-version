package ginservice

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"vw_gateway/internal/conf"
	"vw_gateway/internal/service/ginservice/service"
)

var ProviderSet = wire.NewSet(
	NewGinEngine,
	gs.NewVideoDownloadFileService,
)

func NewGinEngine(jwt *conf.JWT, redis *redis.ClusterClient) *gin.Engine {
	r := gin.Default()
	r.Use(
		cors.New(defineCorsConfig()),
		kgin.Middlewares(
			recovery.Recovery(),
			tracing.Server(),
			//TODO: 添加 jwt 认证中间件，需要设置白名单。目前暂不清楚如何为 gin 添加中间件路由白名单，故暂不添加 jwt 鉴权中间件。
			//auth.JwtAuth(jwt.Secret, time.Duration(jwt.AccessTokenExpiration)*time.Hour, redis),
		))

	return r
}

func defineCorsConfig() cors.Config {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowWebSockets = true
	c.AddAllowHeaders("Authorization")
	return c
}
