package types

import "CampusRecruitment/pkg/models"

type CreateJobForm struct {
	baseForm
	JobName   string    `json:"jobName" form:"jobName" binding:"required"`       //职位名称
	PublishId models.Id `json:"publishId" form:"publishId" swaggerignore:"true"` //发布用户id
	CompId    models.Id `json:"compId" form:"compId" swaggerignore:"true"`       //公司id
	MinWage   int       `json:"minWage" form:"minWage" binding:"required"`       //最小薪资
	MaxWage   int       `json:"maxWage" form:"maxWage" binding:"required"`       //最大薪资
	JobNum    string    `json:"jobNum" form:"jobNum" binding:"required"`         //岗位数量
	Desc      string    `json:"desc" form:"desc" binding:"required"`             //工作描述
	City      string    `json:"city" form:"city" binding:"required"`             //工作城市
	Address   string    `json:"address" form:"address" binding:"required"`       //工作地址
	Tags      []string  `json:"tags" form:"tags"`                                //标签
	State     string    `json:"state" form:"state"`                              //工作状态
}

type QueryJobsForm struct {
	pageForm
	Q string `json:"q" form:"q"`
}

type SearchJobsCondForm struct {
	pageForm
	City       string `json:"city" form:"city"`
	Wage       string `json:"wage" form:"wage"'` //工资区间 'unlimited','3k-','3-5k','5-10k','10-15k','15-20k','20k+'
	CompPeople string `json:"compPeople" form:"compPeople"`
	CompType   string `json:"compType" form:"compType"`
}

type CloseJobsForm struct {
	baseForm
	Ids []models.Id `json:"ids" form:"ids"`
}

type JobCondForm struct {
	baseForm
	City        string `json:"city" form:"city"`
	WageSection string `json:"wageSection" form:"wageSection"'` //工资区间 'unlimited','3k-','3-5k','5-10k','10-15k','15-20k','20k+'
}

type GetJobDetailsForm struct {
	baseForm
	JobId     models.Id `uri:"id" binding:"required" swaggerignore:"true"`     //job id
	PublishId models.Id `json:"publishId" form:"publishId" binding:"required"` //发布者id
	CompId    models.Id `json:"compId" form:"compId" binding:"required"`       //发布公司id
}
