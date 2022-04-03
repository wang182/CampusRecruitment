package ctx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	perrors "github.com/pkg/errors"
	"gorm.io/gorm"

	"CampusRecruitment/pkg/config"
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/infra/db"
	"CampusRecruitment/pkg/infra/log"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
)

type RequestContext interface {
	context.Context

	Logger() log.Logger
	Url() *url.URL
	Token() string
	Bind(form types.Form) error
	Response(result interface{}, err error)
}

type Context struct {
	RequestContext

	m sync.Map
}

const (
	MysqlDuplicate = 1062
)

// 将 err 强制转换为 errors.ErrorCode 类型
func ConvertError(err error) errors.ErrorCode {
	if errCode, ok := err.(errors.ErrorCode); ok {
		return errCode
	}

	switch {
	case perrors.Is(err, gorm.ErrRecordNotFound):
		return errors.ErrNotFound
	case os.IsNotExist(err):
		return errors.ErrFileNotExists
	}

	if _, ok := err.(*mysql.MySQLError); ok {
		return errors.AutoDbErr(err)
	}

	switch err.(type) {
	case validator.ValidationErrors:
		return errors.ErrInvalidParams.WithCause(err)
	case *json.SyntaxError:
		return errors.ErrInvalidJSON.WithCause(err)
	}

	return errors.ErrUnknown.WithCause(err)
}

func New(ctx interface{}) *Context {
	switch v := ctx.(type) {
	case *gin.Context:
		return &Context{RequestContext: newGinRequestCtx(v)}
	default:
		panic(fmt.Errorf("unknown context type: %T", ctx))
	}
}

func (ctx *Context) get(key string) (interface{}, bool) {
	val, ok := ctx.m.Load(key)
	return val, ok
}

func (ctx *Context) set(key string, val interface{}) {
	ctx.m.Store(key, val)
}

func (ctx *Context) User() (user types.UserInfo, err error) {
	if v, ok := ctx.get("_userInfo"); !ok {
		claims := types.UserTokenClaims{}
		_, err = jwt.ParseWithClaims(ctx.Token(), &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get().SecretKey), nil
		})
		if err != nil {
			return user, err
		}

		user = types.UserInfo{
			Id:       claims.UserId,
			Username: claims.UserName,
		}
		ctx.set("_userInfo", user)
		return user, nil
	} else {
		return v.(types.UserInfo), nil
	}
}

func (ctx *Context) MustUser() types.UserInfo {
	user, err := ctx.User()
	if err != nil {
		panic(err)
	}
	return user
}

// 获取请求的 host(含 scheme), eg. http://example.org
func (ctx *Context) RequestHost() string {
	u := ctx.RequestContext.Url()
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}

func (ctx *Context) DB() *gorm.DB {
	return db.Get()
}

func (ctx *Context) PageResult(query *gorm.DB, page, pageSize int, destSlice interface{}) (*types.APIPageResult, error) {
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}

	if pageSize > consts.MaxPageSize {
		pageSize = consts.MaxPageSize
	}

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Scan(destSlice).Error; err != nil {
		return nil, err
	}

	return &types.APIPageResult{
		Page:     page,
		PageSize: pageSize,
		Total:    int(count),
		List:     destSlice,
	}, nil
}
