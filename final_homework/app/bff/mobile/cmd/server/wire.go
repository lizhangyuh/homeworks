//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/biz"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/conf"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/data"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/server"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/service"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Service, *conf.Registry, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
