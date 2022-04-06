package types

import "CampusRecruitment/pkg/models"

type CreateJobForm struct {
	baseForm
	JobName   string    `json:"jobName" form:"jobName"`
	PublishId models.Id `json:"publishId" form:"publishId" swaggerignore:"true"`
	CompId    models.Id `json:"compId" form:"compId" swaggerignore:"true"`
	MinWage   int       `json:"minWage" form:"minWage"`
	MaxWage   int       `json:"maxWage" form:"maxWage"`
	JobNum    int       `json:"jobNum" form:"jobNum"`
	Desc      string    `json:"desc" form:"desc"`
	Address   string    `json:"address" form:"address"`
	Tags      []string  `json:"tags" form:"tags"`
	State     string    `json:"state" form:"state"`
}
