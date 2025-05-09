// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"vw_video/internal/biz"
	"vw_video/internal/conf"
	"vw_video/internal/data"
	"vw_video/internal/server"
	"vw_video/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, confService *conf.Service, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMySQL(confData)
	clusterClient := data.NewRedisClusterClient(confData)
	mongoDB := data.NewMongo(confData)
	discovery := data.NewDiscovery(registry)
	userinfoClient := data.NewUserinfoClient(discovery, confService)
	dataData, cleanup, err := data.NewData(db, clusterClient, mongoDB, userinfoClient, logger)
	if err != nil {
		return nil, nil, err
	}
	videoInfoRepo := data.NewVideoInfoRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	videoInfoUsecase := biz.NewVideoInfoUsecase(videoInfoRepo, transaction, logger)
	videoInfoService := service.NewVideoInfo(videoInfoUsecase, logger)
	interactRepo := data.NewInteractRepo(dataData, logger)
	interactUsecase := biz.NewInteractUseCase(interactRepo, videoInfoRepo, transaction, logger)
	interactService := service.NewInteractService(interactUsecase, logger)
	videoCommentRepo := data.NewVideoCommentRepo(dataData, logger)
	videoCommentUsecase := biz.NewVideoCommentUsecase(videoCommentRepo, transaction, logger)
	videoCommentService := service.NewVideoCommentService(videoCommentUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, videoInfoService, interactService, videoCommentService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
