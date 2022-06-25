package service

import (
	"context"

	pb "nvm.com/mrc-api-go/api/bff/mobile/v1"
)

func (s *MoibleService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	return s.auc.SignUp(ctx, req)
}

func (s *MoibleService) SignInByPwd(ctx context.Context, req *pb.SignInByPwdRequest) (*pb.SignInResponse, error) {
	return s.auc.SignInByPwd(ctx, req)
}

func (s *MoibleService) CurrentUser(ctx context.Context, req *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
	return s.uc.CurrentUser(ctx, req)
}
