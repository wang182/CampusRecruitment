package handler

import (
	"CampusRecruitment/pkg/apps"
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"bytes"
	"github.com/gin-gonic/gin"
	"mime"
	"net/http"
	"path"
)

// IconFileUpload 图标上传
// @Tags 图标
// @Summary 图标上传
// @Accept */*
// @Produce */*
// @Param icon formData file true "图标"
// @router /icons [POST]
// @Success 200
func IconFileUpload(c *gin.Context) {
	ac := ctx.New(c)
	file, err := c.FormFile("icon")
	if err != nil {
		ac.Response(nil, err)
	} else {
		ac.Response(apps.IconFileUpload(file))
	}
}

// ReadIconFile 图标下载
// @Tags 图标
// @Summary 图标下载
// @Accept json
// @Produce json
// @Param form query types.IconForm true "parameter"
// @router /icons [GET]
// @Success 200
func ReadIconFile(c *gin.Context) {
	ac := ctx.New(c)
	iconForm := types.IconForm{}
	err := ac.Bind(&iconForm)
	if err != nil {
		return
	}

	buffer := bytes.NewBuffer(nil)
	err = apps.ReadIconFile(ac, buffer, iconForm.Path)
	if err != nil {
		c.String(http.StatusInternalServerError, errors.AutoErrCode(err).Error())
	}
	data := buffer.Bytes()

	contentType := mime.TypeByExtension(path.Ext(iconForm.Path))
	if contentType == "" {
		contentType = http.DetectContentType(data)
	}

	c.Header("Content-Type", contentType)
	_, _ = c.Writer.Write(data)
}
