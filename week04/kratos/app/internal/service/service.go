package service

import (
	v1 "app/api/app/v1"
	"app/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAppService)

type AppService struct {
	v1.UnimplementedAppServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewAppService new a shop service.
func NewAppService(uc *biz.UserUsecase, logger log.Logger) *AppService {
	return &AppService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/app")),
	}
}
