package data

import (
	"context"
	"crypto/sha512"
	"fmt"

	"nvm.com/mrc-api-go/app/user/service/internal/biz"
	"nvm.com/mrc-api-go/app/user/service/internal/pkg/util"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var options = &password.Options{
	SaltLen:      16,
	Iterations:   10000,
	KeyLen:       32,
	HashFunction: sha512.New,
}

type User struct {
	ID       int64
	Nickname string
	Phone    string
	Password string
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	// var user User
	// 检查是否已注册
	result := ur.data.db.Where(&biz.User{Phone: user.Phone}).First(&user)
	if result.RowsAffected >= 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	ph, err := util.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = ph

	res := ur.data.db.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}

	return user, nil

}

func (ur *userRepo) CheckPassword(ctx context.Context, phone, password string) (bool, error) {
	var user User
	res := ur.data.db.Where(&biz.User{Phone: phone}).First(&user)

	if res.Error != nil {
		return false, res.Error
	}

	ok := util.CheckPasswordHash(password, user.Password)
	return ok, nil
}

// UserByPhone .
func (ur *userRepo) UserByPhone(ctx context.Context, phone string) (*biz.User, error) {
	var user User
	result := ur.data.db.Where(&User{Phone: phone}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	if result.Error != nil {
		return nil, errors.New(500, "FIND_USER_ERROR", "find user error")
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	return &biz.User{
		ID:       user.ID,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Password: user.Password,
	}, nil
}

// UserById .
func (ur *userRepo) UserById(ctx context.Context, id int64) (*biz.User, error) {
	var user User
	result := ur.data.db.Where(&User{ID: id}).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	if result.Error != nil {
		return nil, errors.New(500, "FIND_USER_ERROR", "find user error")
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	return &biz.User{
		ID:       user.ID,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Password: user.Password,
	}, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func encrypt(pwd string) string {

	salt, encodedPwd := password.Encode(pwd, options)

	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
}
