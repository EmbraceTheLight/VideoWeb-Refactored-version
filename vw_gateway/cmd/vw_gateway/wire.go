//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"vw_gateway/internal/biz/userbiz"
	"vw_gateway/internal/conf"
	"vw_gateway/internal/data/userdata"
	"vw_gateway/internal/pkg/captcha"
	"vw_gateway/internal/pkg/middlewares"
	"vw_gateway/internal/server"
	"vw_gateway/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, *conf.JWT, *conf.Email, *conf.Trace, *conf.Service, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, userdata.ProviderSet, userbiz.ProviderSet, middlewares.ProviderSet, captcha.Provider, service.ProviderSet, newApp))
}
