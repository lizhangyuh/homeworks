package service

import (
	"context"

	pb "user/api/user/v1"
	"user/internal/biz"

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
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "get user")
	defer span.End()
	user, err := u.uc.UserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	return &pb.UserInfoReply{
		Id:       user.ID,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Password: user.Password,
	}, nil
}

func (s *UserService) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordReply, error) {
	tr := otel.Tracer("service")
	ctx, span := tr.Start(ctx, "check user password")
	defer span.End()

	result, err := s.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)

	if err != nil {
		return nil, err
	}

	return &pb.CheckPasswordReply{
		Success: result,
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
