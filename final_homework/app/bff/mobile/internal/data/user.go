package data

import (
	"context"

	userService "nvm.com/mrc-api-go/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (u *userRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	createUser, err := u.data.uc.CreateUser(c, &userService.CreateUserRequest{
		Nickname: user.Nickname,
		Password: user.Password,
		Phone:    user.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       createUser.Id,
		Phone:    createUser.Phone,
		Nickname: createUser.Nickname,
	}, nil
}

func (u *userRepo) CheckPassword(c context.Context, phone string, password string) (bool, error) {
	check, err := u.data.uc.CheckPassword(c, &userService.CheckPasswordRequest{
		Phone:    phone,
		Password: password,
	})

	if err != nil {
		return false, err
	} else {
		return check.Ok, nil
	}
}

func (u *userRepo) UserByPhone(c context.Context, phone string) (*biz.User, error) {
	user, err := u.data.uc.GetUserByPhone(c, &userService.PhoneRequest{Phone: phone})

	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:       user.Id,
		Phone:    user.Phone,
		Nickname: user.Nickname,
	}, nil
}

func (u *userRepo) UserById(cxt context.Context, id int64) (*biz.User, error) {
	user, err := u.data.uc.GetUserById(cxt, &userService.IdRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:       user.Id,
		Phone:    user.Phone,
		Nickname: user.Nickname,
	}, nil
}
