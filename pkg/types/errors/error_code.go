package errors

import (
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ErrorCode struct {
	code    int
	message string
	cause   error

	statusCode int
}

var errorCodes = make(map[int]struct{})

// 基于 err 自动获取 errorCode
// 如果 err 为 ErrorCode 类型，则直接返回，
// 否则，将 err 设置为 defaultCode 的 cause，然后返回
func AutoErrCode(err error, defaultCodes ...ErrorCode) ErrorCode {
	if errCode, ok := err.(ErrorCode); ok {
		return errCode
	} else {
		if len(defaultCodes) == 0 {
			return ErrUnknown.WithCause(err)
		}
		return defaultCodes[0].WithCause(err)
	}
}

// Wrap 替代 errors.Warp()
// 如果传入的是 ErrCode 类型，该函数返回的也是 ErrCode
func Wrap(err error, message string) error {
	return Wrapf(err, message)
}

// Wrapf 同 Warp
func Wrapf(err error, format string, args ...interface{}) error {
	if errCode, ok := err.(ErrorCode); ok {
		if cause := errCode.Cause(); cause != nil {
			return errCode.WithCause(errors.Wrapf(cause, format, args...))
		} else {
			return errCode.WithCause(fmt.Errorf(format, args...))
		}
	} else {
		return errors.Wrapf(err, format, args...)
	}
}

const (
	MysqlDuplicate = 1062
)

func AutoDbErr(err error) ErrorCode {
	if errCode, ok := err.(ErrorCode); ok {
		return errCode
	}

	if myErr, ok := err.(*mysql.MySQLError); ok {
		switch myErr.Number {
		case MysqlDuplicate:
			return ErrAlreadyExists.WithCause(err)
		}
	}
	return ErrDBError.WithCause(err)
}

func newErrorCode(message string, code int) ErrorCode {
	if _, ok := errorCodes[code]; ok {
		panic(fmt.Errorf("duplicate error code: %v", code))
	} else {
		errorCodes[code] = struct{}{}
	}
	return ErrorCode{message: message, code: code}
}

func (e ErrorCode) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %s", e.message, e.cause.Error())
	} else {
		return e.message
	}
}

func (e ErrorCode) Code() int {
	return e.code
}

func (e ErrorCode) Message() string {
	return e.message
}

func (e ErrorCode) WithCause(err error) ErrorCode {
	e.cause = err
	return e
}

func (e ErrorCode) Cause() error {
	return e.cause
}

func (e ErrorCode) WithStatus(statusCode int) ErrorCode {
	e.statusCode = statusCode
	return e
}

func (e ErrorCode) Status() int {
	if e.statusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.statusCode
}

var (
	notFoundErrCodePrefix = ErrNotFound.code / 1000
)

func IsNotFoundErr(err error) bool {
	if err == nil {
		return false
	}

	if e, ok := err.(ErrorCode); ok && e.code/1000 == notFoundErrCodePrefix {
		return true
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
