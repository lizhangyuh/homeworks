// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnknownError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == BffMobileUserErrorReason_UNKNOWN_ERROR.String() && e.Code == 500
}

func ErrorUnknownError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, BffMobileUserErrorReason_UNKNOWN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsLoginFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == BffMobileUserErrorReason_LOGIN_FAILED.String() && e.Code == 401
}

func ErrorLoginFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(401, BffMobileUserErrorReason_LOGIN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsUsernameConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == BffMobileUserErrorReason_USERNAME_CONFLICT.String() && e.Code == 409
}

func ErrorUsernameConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, BffMobileUserErrorReason_USERNAME_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsRegisterFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == BffMobileUserErrorReason_REGISTER_FAILED.String() && e.Code == 541
}

func ErrorRegisterFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(541, BffMobileUserErrorReason_REGISTER_FAILED.String(), fmt.Sprintf(format, args...))
}
