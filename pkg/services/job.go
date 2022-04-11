package services

import (
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"CampusRecruitment/pkg/types/resps"
	"gorm.io/gorm"
	"strings"
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
		City:        form.City,
		Address:     form.Address,
		Tags:        ChangeArrayToString(form.Tags),
	}
	if err := db.Model(&models.Job{}).Create(&job).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return GetJobById(db, job.Id)

}

func QueryJobs(db *gorm.DB, q string) *gorm.DB {
	return db.Model(&models.Job{}).Where("job_name LIKE ?", "%"+q+"%").Order("created_at")
}

func SearchJobsWithCond(db *gorm.DB, cond *types.JobCondForm) *gorm.DB {
	return db.Table(models.Job{}.TableName()).Where("state", "active").Where(cond)
}

func GetJobsIdByCompId(db *gorm.DB, id models.Id) ([]models.Id, error) {
	ids := make([]models.Id, 0)
	if err := db.Model(&models.Job{}).Where("comp_id", id).Select("id").Scan(&ids).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return ids, nil
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

func DeleteJob(db *gorm.DB, id models.Id) error {
	job := models.Job{}
	if err := db.Model(&models.Job{}).Where("id", id).Delete(&job).Error; err != nil {
		return errors.AutoDbErr(err)
	}
	return nil
}

func CLoseJob(db *gorm.DB, id models.Id) error {
	ids := make([]models.Id, 0)
	ids[0] = id
	if err := UpdateJobsState(db, ids, "inactive"); err != nil {
		return err
	}
	return nil
}

func UpdateJobsState(db *gorm.DB, ids []models.Id, state string) error {
	if err := db.Model(&models.Job{}).Where("id IN ?", ids).Update("state", state).Error; err != nil {
		return errors.AutoDbErr(err)
	}
	return nil
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

func ChangeArrayToString(arrays []string) string {
	str := strings.Join(arrays, ",")
	return str
}

func ChangeStringToArray(str string) []string {
	arr := strings.Split(str, ",")
	return arr
}

func GetHotJobsByCompId(db *gorm.DB, compId models.Id) ([]resps.HotJobResp, error) {
	hotJobs := make([]resps.HotJobResp, 0)
	if err := db.Model(&models.Job{}).Where("comp_id", compId).Order("updated_at").Limit(6).Scan(&hotJobs).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return hotJobs, nil
}

func GetJobNumByCompId(db *gorm.DB, compId models.Id) int {
	var num int
	db.Model(&models.Job{}).Raw("SELECT COUNT(1) FROM t_job WHERE comp_id = ?", compId).Scan(&num)
	return num
}
