package biz

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	v1 "nvm.com/mrc-api-go/api/bff/mobile/v1"

	"nvm.com/mrc-api-go/app/bff/mobile/internal/conf"
)

// 定义错误信息
var (
	ErrLoginFailed     = errors.New("login failed")
	ErrUserNotFound    = errors.New("user not found")
	ErrPhoneUsed       = errors.New("phone is already used")
	ErrPasswordInvalid = errors.New("password invalid")
)

type AuthUseCase struct {
	key  string
	repo UserRepo
}

func NewAuthUseCase(conf *conf.Auth, ur UserRepo) *AuthUseCase {
	return &AuthUseCase{
		key:  conf.ApiKey,
		repo: ur,
	}
}

func (uc *AuthUseCase) SignInByPwd(ctx context.Context, req *v1.SignInByPwdRequest) (*v1.SignInResponse, error) {

	// check user
	user, err := uc.repo.UserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// check password
	check, err := uc.repo.CheckPassword(ctx, req.Phone, req.Password)

	if err != nil {
		return nil, ErrPasswordInvalid
	} else {

		if !check {
			return nil, ErrLoginFailed
		}

		// generate token
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.Id,
		})

		token, err := claims.SignedString([]byte(uc.key))
		if err != nil {
			return nil, err
		}

		return &v1.SignInResponse{
			Token: fmt.Sprintf("Bearer %s", token),
		}, nil
	}

}

func (uc *AuthUseCase) SignUp(ctx context.Context, req *v1.SignUpRequest) (*v1.SignUpResponse, error) {

	// check phone
	_, err := uc.repo.UserByPhone(ctx, req.Phone)
	if errors.Is(err, ErrUserNotFound) {
		return nil, v1.ErrorRegisterFailed("phone is already used")
	}

	// create user
	newUser := &User{Phone: req.Phone, Nickname: req.Nickname, Password: req.Password}
	user, err := uc.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, v1.ErrorRegisterFailed("create user failed: %s", err.Error())
	}

	return &v1.SignUpResponse{
		Id:       user.Id,
		Nickname: user.Nickname,
		Phone:    user.Phone,
	}, nil
}
