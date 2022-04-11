package types

import "CampusRecruitment/pkg/models"

type CompRegisterForm struct {
	baseForm
	CompName     string `json:"compName" form:"compName" binding:"required"`         //公司名字
	Logo         string `json:"logo" form:"logo" binding:"required"`                 //公司logo
	CompType     string `json:"compType" form:"compType" binding:"required"`         //公司类型 'mall','game','medical','hardware','software','network','finance','video','education','other'
	PeopleNum    string `json:"peopleNum" form:"peopleNum" binding:"required"`       //公司人数区间 '20','99','500','1000','9999','10000'
	City         string `json:"city" form:"city" binding:"required"`                 //公司所在城市
	Introduction string `json:"introduction" form:"introduction" binding:"required"` //公司简介
	Address      string `json:"address" form:"address" binding:"required"`           //公司详细地址
	Url          string `json:"url" form:"url" binding:""`                           //公司官网
}

type QueryCompWithCondForm struct {
	pageForm
	City      string `json:"place"form:"place"`
	PeopleNum string `json:"peopleNum" form:"peopleNum"`
	CompType  string `json:"compType" form:"compType"`
}

type QueryCompForm struct {
	pageForm
	Q string `json:"q" form:"q"`
}

type UpdateCompForm struct {
	baseForm
	Id           models.Id `uri:"id" json:"id" swaggerignore:"true" binding:"required"`
	CompName     string    `json:"compName" form:"compName" binding:""`
	Logo         string    `json:"logo" form:"logo" binding:""`
	CompType     string    `json:"compType" form:"compType" binding:""`
	PeopleNum    string    `json:"peopleNum" form:"peopleNum" binding:""`
	City         string    `json:"city" form:"city" binding:""`
	Introduction string    `json:"introduction" form:"introduction" binding:""`
	Address      string    `json:"address" form:"address" binding:""`
	Url          string    `json:"url" form:"url" binding:""`
	State        string    `json:"state" form:"url" binding:""`
}
