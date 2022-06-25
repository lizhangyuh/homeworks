package biz

import (
	"context"
	"errors"

	v1 "nvm.com/mrc-api-go/api/bff/mobile/v1"
	"nvm.com/mrc-api-go/pkg/auth"

	"github.com/go-kratos/kratos/v2/log"
)

// 定义错误信息
var (
	ErrCaptchaInvalid = errors.New("verification code error")
	ErrAuthFailed     = errors.New("authentication failed")
)

type User struct {
	Id       int64
	Phone    string
	Nickname string
	Password string
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	UserByPhone(cxt context.Context, phone string) (*User, error)
	UserById(cxt context.Context, id int64) (*User, error)
	CheckPassword(cxt context.Context, pwd string, encryptedPwd string) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/client"))}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	newUser := &User{Phone: req.Phone, Nickname: req.Nickname, Password: req.Password}
	createdUser, err := uc.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &v1.SignUpResponse{
		Id:       createdUser.Id,
		Phone:    createdUser.Phone,
		Nickname: createdUser.Nickname,
	}, nil
}

func (uc *UserUsecase) CurrentUser(ctx context.Context, req *v1.CurrentUserRequest) (*v1.CurrentUserResponse, error) {

	userId, err := auth.CurrentUserId(ctx)

	if err != nil {
		return nil, err
	}

	user, err := uc.repo.UserById(ctx, userId.(int64))

	if err != nil {
		return nil, err
	}

	return &v1.CurrentUserResponse{
		Id:       user.Id,
		Phone:    user.Phone,
		Nickname: user.Nickname,
	}, nil
}
