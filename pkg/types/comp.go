package types

import "CampusRecruitment/pkg/models"

type CompRegisterForm struct {
	baseForm
	CompName     string `json:"compName" form:"compName" binding:"required"`
	Logo         string `json:"logo" form:"logo" binding:"required"`
	CompType     string `json:"compType" form:"compType" binding:"required"`
	PeopleNum    string `json:"peopleNum" form:"peopleNum" binding:"required"`
	City         string `json:"city" form:"city" binding:"required"`
	Introduction string `json:"introduction" form:"introduction" binding:"required"`
	Address      string `json:"address" form:"address" binding:"required"`
	Url          string `json:"url" form:"url" binding:""`
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
	CompId       models.Id `uri:"compId" json:"compId" swaggerignore:"true" binding:"required"`
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
