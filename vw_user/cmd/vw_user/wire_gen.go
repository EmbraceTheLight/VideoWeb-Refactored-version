// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"vw_user/internal/biz"
	"vw_user/internal/conf"
	"vw_user/internal/data"
	"vw_user/internal/server"
	"vw_user/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMySQL(confData)
	clusterClient := data.NewRedisClusterClient(confData)
	client := data.NewMongo(confData)
	dataData, cleanup, err := data.NewData(db, clusterClient, client, logger)
	if err != nil {
		return nil, nil, err
	}
	userIdentityRepo := data.NewUserRepo(dataData, logger)
	userInfoRepo := data.NewUserInfoRepo(dataData, logger)
	userIdentityUsecase := biz.NewUserIdentityUsecase(userIdentityRepo, userInfoRepo, logger)
	userIdentityService := service.NewUserIdentityService(userIdentityUsecase, logger)
	userInfoUsecase := biz.NewUserInfoUsecase(userInfoRepo, logger)
	userInfoService := service.NewUserInfoService(userInfoUsecase, logger)
	captchaRepo := data.NewCaptRepo(dataData, logger)
	captchaUsecase := biz.NewCaptchaUsecase(logger, captchaRepo)
	captchaService := service.NewCaptchaService(logger, captchaUsecase)
	fileUsecase := biz.NewFileUsecase(logger)
	fileService := service.NewFileService(logger, fileUsecase)
	grpcServer := server.NewGRPCServer(confServer, userIdentityService, userInfoService, captchaService, fileService, logger)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
