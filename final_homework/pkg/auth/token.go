package auth

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

var (
	ErrAuthFailed = errors.New("authentication failed")
)

func CurrentUserId(ctx context.Context) (interface{}, error) {
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwtv4.MapClaims)
		if c["user_id"] == nil {
			return nil, ErrAuthFailed
		}

		userId := int64(c["user_id"].(float64))

		return userId, nil
	} else {
		return nil, ErrAuthFailed
	}
}
