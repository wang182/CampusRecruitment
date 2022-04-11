package handler

import (
	"CampusRecruitment/pkg/apps"
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/types"
	"github.com/gin-gonic/gin"
)

// CreateJob 发布职位
// @Tags Job
// @Summary 发布职位
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json body types.CreateJobForm true "parameter"
// @router /job [POST]
// @Success 200 {object} types.APIResponse{result=models.Job}
func CreateJob(c *gin.Context) {
	ac := ctx.New(c)
	form := types.CreateJobForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.CreateJob(ac, &form))
}

// QueryJobs 查询工作（模糊查询）
// @Tags Job
// @Summary 查询职位
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json query types.QueryJobsForm true "parameter"
// @router /jobs [GET]
// @Success 200 {object} types.APIResponse{result=types.APIPageResult{list=[]models.Job}}
func QueryJobs(c *gin.Context) {
	ac := ctx.New(c)
	form := types.QueryJobsForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.QueryJobs(ac, &form))
}

// AdminQueryJobs 查询工作（后台管理接口）
// @Tags Job
// @Summary 查询工作（后台管理接口）
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json query types.QueryJobsForm true "parameter"
// @router /admin/jobs [GET]
// @Success 200 {object} types.APIResponse{result=types.APIPageResult{list=[]models.Job}}
func AdminQueryJobs(c *gin.Context) {
	ac := ctx.New(c)
	form := types.QueryJobsForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.AdminQueryJobs(ac, &form))
}

// SearchJobsWithCond 查询职位(条件查询)
// @Tags Job
// @Summary 查询职位
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json query types.SearchJobsCondForm true "parameter"
// @router /jobs/cond [GET]
// @Success 200 {object} types.APIResponse{result=types.APIPageResult{list=[]models.Job}}
func SearchJobsWithCond(c *gin.Context) {
	ac := ctx.New(c)
	form := types.SearchJobsCondForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.QueryJobsWithCond(ac, &form))
}

// DeleteJob 删除工作（管理员）
// @Tags Job
// @Summary 删除工作
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path string true "job id"
// @router /jobs/{id} [DELETE]
// @Success 200 {object} types.APIResponse{}
func DeleteJob(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(apps.DeleteJob(ac, models.Id(c.Param("id"))))
}

// CloseJob 关闭岗位
// @Tags Job
// @Summary 关闭岗位
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path string true "job id"
// @router /jobs/{id}/close [PUT]
// @Success 200 {object} types.APIResponse{}
func CloseJob(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(apps.CloseJob(ac, models.Id(c.Param("id"))))
}

// CloseJobsByIds 批量关闭职位
// @Tags Job
// @Summary 批量关闭职位
// @Accept json
// @Produce json
// @Security AuthToken
// @Param form body types.CloseJobsForm true "parameter"
// @router /jobs/{id}/close [PUT]
// @Success 200 {object} types.APIResponse{}
func CloseJobsByIds(c *gin.Context) {
	ac := ctx.New(c)
	form := types.CloseJobsForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.CloseJobsByIds(ac, &form))
}

// GetJobDetails 获取职位详情
// @Tags Job
// @Summary 获取职位详情
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path string true "job id"
// @Param form query types.GetJobDetailsForm true "parameter"
// @router /jobs/{id}/details [GET]
// @Success 200 {object} types.APIResponse{}
func GetJobDetails(c *gin.Context) {
	ac := ctx.New(c)
	form := types.GetJobDetailsForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.GetJobDetails(ac, &form))
}

// ApproveJob 岗位通过申请
// @Tags Job
// @Summary 岗位通过申请
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path string true "job id"
// @router /jobs/{id} [PUT]
// @Success 200 {object} types.APIResponse{}
func ApproveJob(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(apps.ApproveJob(ac, models.Id(c.Param("id"))))
}
