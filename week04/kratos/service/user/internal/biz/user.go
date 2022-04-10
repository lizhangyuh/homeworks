package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
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

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) UserByPhone(cxt context.Context, phone string) (*User, error) {
	return uc.repo.UserByPhone(cxt, phone)
}

func (uc *UserUseCase) CheckPassword(ctx context.Context, pwd, encryptedPwd string) (bool, error) {
	return uc.repo.CheckPassword(ctx, pwd, encryptedPwd)
}
