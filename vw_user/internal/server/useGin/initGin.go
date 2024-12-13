package useGin

import (
	"context"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"time"
	"vw_user/internal/conf"
	"vw_user/internal/pkg/middlewares/auth"
)

func NewWhitelistMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		return true
	}
}

func NewGinEngine(jwt *conf.JWT) *gin.Engine {
	r := gin.Default()
	r.Use(kgin.Middlewares(
		recovery.Recovery(),
		tracing.Server(),
		// jwt鉴权中间件
		selector.Server(
			auth.JwtAuth(jwt.Secret, time.Duration(jwt.AccessTokenExpiration)*time.Hour),
		).Match(NewWhitelistMatcher()).Build(),
	))
	return r
}
