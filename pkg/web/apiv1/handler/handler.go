package handler

import (
	"github.com/gin-gonic/gin"

	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/types/errors"
)

func NotImplemented(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(nil, errors.ErrNotImplemented)
}
