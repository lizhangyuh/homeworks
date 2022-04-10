package data

import (
	userService "app/api/service/user/v1"
	"app/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
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
		ID:       createUser.Id,
		Phone:    createUser.Phone,
		Nickname: createUser.Nickname,
	}, nil
}

func (u *userRepo) CheckPassword(c context.Context, pwd string, encryptedPwd string) (bool, error) {
	check, err := u.data.uc.CheckPassword(c, &userService.CheckPasswordRequest{
		Password:          pwd,
		EncryptedPassword: encryptedPwd,
	})

	if err != nil {
		return false, err
	} else {
		return check.Success, nil
	}
}

func (u *userRepo) UserByPhone(c context.Context, phone string) (*biz.User, error) {
	user, err := u.data.uc.GetUserByPhone(c, &userService.PhoneRequest{Phone: phone})

	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:       user.Id,
		Phone:    user.Phone,
		Nickname: user.Nickname,
		Password: user.Password,
	}, nil
}
