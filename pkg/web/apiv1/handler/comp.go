package handler

import (
	"CampusRecruitment/pkg/apps"
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/types"
	"github.com/gin-gonic/gin"
)

// CompRegister 公司注册
// @Tags Comp
// @Summary 公司注册
// @Accept json
// @Produce json
// @Param json body types.CompRegisterForm true "parameter"
// @router /comp [POST]
// @Success 200 {object} types.APIResponse{result=models.Comp}
func CompRegister(c *gin.Context) {
	ac := ctx.New(c)
	form := types.CompRegisterForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.CompRegister(ac, &form))
}

// QueryCompWithCond 条件筛选公司
// @Tags Comp
// @Summary 条件筛选公司
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json query types.QueryCompWithCondForm true "parameter"
// @router /comp/cond [GET]
// @Success 200 {object} types.APIResponse{result=types.APIPageResult{list=[]resps.CompCondResp}}
func QueryCompWithCond(c *gin.Context) {
	ac := ctx.New(c)
	form := types.QueryCompWithCondForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.QueryCompWithCond(ac, &form))
}

// QueryComp 查询所有公司（支持模糊查询）
// @Tags Comp
// @Summary 查询所有公司
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json query types.QueryCompForm true "parameter"
// @router /comp [GET]
// @Success 200 {object} types.APIResponse{result=types.APIPageResult{list=[]resps.CompCondResp}}
func QueryComp(c *gin.Context) {
	ac := ctx.New(c)
	form := types.QueryCompForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.QueryComp(ac, &form))
}

// CloseComp 关闭公司（管理）
// @Tags Comp
// @Summary 关闭公司
// @Accept json
// @Produce json
// @Security AuthToken
// @Param compId path string true "comp id"
// @router /comp/{compId}/close [PUT]
// @Success 200 {object} types.APIResponse{}
func CloseComp(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(apps.CloseComp(ac, models.Id(c.Param("id"))))
}

// ApproveComp 更新公司状态 （管理）
// @Tags Comp
// @Summary 更新公司状态
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path string true "comp id"
// @router /comp/{id}/approve [PUT]
// @Success 200 {object} types.APIResponse{}
func ApproveComp(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(apps.ApproveComp(ac, models.Id(c.Param("id"))))
}

// UpdateComp 更新公司信息
// @Tags Comp
// @Summary 更新公司信息
// @Accept json
// @Produce json
// @Security AuthToken
// @Param compId path string true "comp id"
// @Param form body types.UpdateCompForm true "parameter"
// @router /comp/{compId} [PUT]
// @Success 200 {object} types.APIResponse{}
func UpdateComp(c *gin.Context) {
	ac := ctx.New(c)
	form := types.UpdateCompForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.UpdateComp(ac, &form))
}
