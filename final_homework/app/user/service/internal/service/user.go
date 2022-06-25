package service

import (
	"context"

	pb "nvm.com/mrc-api-go/api/user/service/v1"
	"nvm.com/mrc-api-go/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
)

type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user, err := s.uc.Create(ctx, &biz.User{
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{
		Id:       user.ID,
		Nickname: user.Nickname,
		Phone:    user.Phone,
	}, nil
}

// GetUserByPhone .
func (u *UserService) GetUserByPhone(ctx context.Context, req *pb.PhoneRequest) (*pb.UserInfoReply, error) {
	user, err := u.uc.UserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	return &pb.UserInfoReply{
		Id:       user.ID,
		Nickname: user.Nickname,
		Phone:    user.Phone,
	}, nil
}

// GetUserById .
func (u *UserService) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.UserInfoReply, error) {
	user, err := u.uc.UserById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserInfoReply{
		Id:       user.ID,
		Nickname: user.Nickname,
		Phone:    user.Phone,
	}, nil
}

func (s *UserService) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordReply, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "check user password")
	defer span.End()

	result, err := s.uc.CheckPassword(ctx, req.Phone, req.Password)

	if err != nil {
		return nil, err
	}

	return &pb.CheckPasswordReply{
		Ok: result,
	}, nil
}

// func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
// 	return &pb.UpdateUserReply{}, nil
// }
// func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
// 	return &pb.DeleteUserReply{}, nil
// }
// func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
// 	return &pb.GetUserReply{}, nil
// }
// func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
// 	return &pb.ListUserReply{}, nil
// }
