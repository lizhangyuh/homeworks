// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/biz"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/conf"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/data"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/server"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, confService *conf.Service, registry *conf.Registry, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	discovery := data.NewDiscovery(registry)
	userClient := data.NewUserServiceClient(auth, confService, discovery, tracerProvider)
	dataData, err := data.NewData(confData, userClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	authUseCase := biz.NewAuthUseCase(auth, userRepo)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	moibleService := service.NewMoibleService(authUseCase, userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, tracerProvider, moibleService, logger)
	registrar := data.NewRegistrar(registry)
	app := newApp(logger, httpServer, registrar)
	return app, func() {
	}, nil
}
