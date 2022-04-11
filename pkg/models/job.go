package models

import "CampusRecruitment/pkg/consts"

type Job struct {
	TimedModel
	JobName     string `gorm:"not null;"`
	PublishId   Id     `gorm:"not null"`
	CompId      Id     `gorm:"not null"`
	WageSection string `gorm:"type:enum('unlimited','3k-','3-5k','5-10k','10-15k','15-20k','20k+');not null"`
	MaxWage     int    `gorm:"not null"`
	MinWage     int    `gorm:"not null"`
	JobNum      string `gorm:"not null"`
	Desc        string `gorm:"type:text;not null;"`
	City        string `gorm:"not null;"`
	Address     string `gorm:"not null"`
	Tags        string `gorm:""`
	State       string `gorm:"type:enum('approving','active','inactive');default:approving;not null"`
}

func (Job) TableName() string {
	return consts.TablePrefix + "job"
}
