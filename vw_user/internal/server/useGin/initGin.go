package useGin

import (
	"context"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

func NewWhitelistMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		return true
	}
}

func NewGinEngine() *gin.Engine {
	r := gin.Default()
	r.Use(kgin.Middlewares(
		recovery.Recovery(),
		tracing.Server(),
	))
	return r
}
