package service

import (
	v1 "nvm.com/mrc-api-go/api/bff/mobile/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMoibleService)

type MoibleService struct {
	v1.UnimplementedMobileServer

	auc *biz.AuthUseCase
	uc  *biz.UserUsecase
	pc  *biz.PostUsecase
	log *log.Helper
}

// NewMoibleService new a shop service.
func NewMoibleService(auc *biz.AuthUseCase, uc *biz.UserUsecase, logger log.Logger) *MoibleService {
	return &MoibleService{
		auc: auc,
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/mobile")),
	}
}
