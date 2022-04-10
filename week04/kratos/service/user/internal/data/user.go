package data

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"user/internal/biz"

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

	user.Password = encrypt(user.Password)

	res := ur.data.db.Create(&user)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}

	return user, nil

}

func (ur *userRepo) CheckPassword(ctx context.Context, pwd, encryptedPwd string) (bool, error) {
	passwordInfo := strings.Split(encryptedPwd, "$")
	result := password.Verify(pwd, passwordInfo[2], passwordInfo[3], options)
	return result, nil
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
