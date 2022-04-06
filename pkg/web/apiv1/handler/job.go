package handler

import (
	"CampusRecruitment/pkg/apps"
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/types"
	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	ac := ctx.New(c)
	form := types.CreateJobForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.CreateJob(ac, &form))
}
