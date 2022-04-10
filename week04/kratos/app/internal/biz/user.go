package biz

import (
	"context"

	v1 "app/api/app/v1"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

// 定义错误信息
var (
	ErrLoginFailed     = errors.New("login failed")
	ErrUserNotFound    = errors.New("user not found")
	ErrPasswordInvalid = errors.New("password invalid")
)

type User struct {
	ID       int64
	Phone    string
	Nickname string
	Password string
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	UserByPhone(cxt context.Context, phone string) (*User, error)
	CheckPassword(cxt context.Context, pwd string, encryptedPwd string) (bool, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/client"))}
}

func (uc *UserUsecase) SignInByPassword(ctx context.Context, req *v1.SignInRequest) (*v1.SignInResponse, error) {

	user, err := uc.repo.UserByPhone(ctx, req.Phone)

	if err != nil {
		return nil, ErrUserNotFound
	}

	check, err := uc.repo.CheckPassword(ctx, req.Password, user.Password)

	if err != nil {
		return nil, ErrPasswordInvalid
	} else {

		if !check {
			return nil, ErrLoginFailed
		}

		return &v1.SignInResponse{
			Id:       user.ID,
			Phone:    user.Phone,
			Nickname: user.Nickname,
		}, nil
	}

}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	newUser := &User{Phone: req.Phone, Nickname: req.Nickname, Password: req.Password}
	createdUser, err := uc.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &v1.SignUpResponse{
		Id:       createdUser.ID,
		Phone:    createdUser.Phone,
		Nickname: createdUser.Nickname,
	}, nil
}
