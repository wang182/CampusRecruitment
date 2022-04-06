package services

import (
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"gorm.io/gorm"
)

func CreateJob(db *gorm.DB, form *types.CreateJobForm) (*models.Job, error) {
	job := models.Job{
		JobName:     form.JobName,
		PublishId:   form.PublishId,
		CompId:      form.CompId,
		MinWage:     form.MinWage,
		MaxWage:     form.MaxWage,
		WageSection: SelectWageSection(form.MinWage),
		JobNum:      form.JobNum,
		Desc:        form.Desc,
		Address:     form.Address,
		Tags:        form.Tags,
	}
	if err := db.Model(&models.Job{}).Create(&job).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return GetJobById(db, job.Id)

}

func GetJobById(db *gorm.DB, id models.Id) (*models.Job, error) {
	job := models.Job{}
	if err := db.Model(&models.Job{}).Where("id", id).First(&job).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.AutoDbErr(err)
	}
	return &job, nil
}

func SelectWageSection(minWage int) string {
	if minWage <= 3000 {
		return consts.WageA
	}
	if minWage <= 5000 && minWage > 3000 {
		return consts.WageB
	}
	if minWage <= 10000 && minWage > 5000 {
		return consts.WageC
	}
	if minWage <= 15000 && minWage > 10000 {
		return consts.WageD
	}
	if minWage <= 20000 && minWage > 15000 {
		return consts.WageE
	}
	if minWage > 20000 {
		return consts.WageF
	}
	return consts.Wage0
}
