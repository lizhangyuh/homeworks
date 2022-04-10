// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"app/internal/biz"
	"app/internal/conf"
	"app/internal/data"
	"app/internal/server"
	"app/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, confService *conf.Service, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(confService, discovery)
	dataData, err := data.NewData(confData, userClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	appService := service.NewAppService(userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, appService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
	}, nil
}