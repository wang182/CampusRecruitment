package ctx

import (
	"CampusRecruitment/pkg/infra/log"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"CampusRecruitment/pkg/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type GinRequestCtx struct {
	*gin.Context

	logger log.Logger
	token  *string
}

func newGinRequestCtx(gctx *gin.Context) *GinRequestCtx {
	traceId := utils.RandomStr(8)
	return &GinRequestCtx{
		Context: gctx,
		logger:  log.Get().WithField("traceId", traceId),
	}
}

func (c *GinRequestCtx) Bind(form types.Form) error {
	if err := c.Context.ShouldBind(form); err != nil {
		c.errorResposne(nil, errors.ErrBadRequest.WithCause(err).WithStatus(http.StatusBadRequest))
		return err
	}
	if err := form.Validate(); err != nil {
		c.errorResposne(nil, errors.ErrBadRequest.WithCause(err).WithStatus(http.StatusBadRequest))
		return err
	}
	return nil
}

func (c *GinRequestCtx) Logger() log.Logger {
	return c.logger
}

func (c *GinRequestCtx) Url() *url.URL {
	u := *c.Request.URL

	u.Host = c.Request.Host
	if host := c.Request.Header.Get("X-Forwarded-Host"); host != "" {
		u.Host = host
	}

	// registry 默认使用 http 协议，如果前端反代(如 nginx) 使用了 https 协议，
	// 则反代时应该在设置 X-Forwarded-Proto 头为 https
	u.Scheme = "http"
	if scheme := c.Request.Header.Get("X-Forwarded-Proto"); scheme != "" {
		u.Scheme = scheme
	}
	return &u
}

func (c *GinRequestCtx) Token() string {
	if c.token == nil {
		auth := c.Context.GetHeader("Authorization")
		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		c.token = &token
	}
	return *c.token
}

func (c *GinRequestCtx) Response(result interface{}, err error) {
	if err != nil {
		c.errorResposne(result, err)
	} else {
		c.Context.JSON(http.StatusOK, &types.APIResponse{
			Code:          0,
			Message:       "",
			MessageDetail: "",
			Result:        result,
		})
	}
}

func (c *GinRequestCtx) errorResposne(result interface{}, err error) {
	errCode := ConvertError(err)
	resp := types.APIResponse{
		Code:    errCode.Code(),
		Message: errCode.Message(),
		Result:  result,
	}
	if errCode.Cause() != nil {
		resp.MessageDetail = errCode.Cause().Error()
	}
	c.logger.WithField("handler", c.HandlerName()).Infof("response error: %+v", resp)
	c.Context.JSON(errCode.Status(), resp)
}
