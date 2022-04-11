package services

import (
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"gorm.io/gorm"
)

func CreateComp(db *gorm.DB, form *types.CompRegisterForm) (*models.Comp, error) {
	comp := models.Comp{
		CompName:     form.CompName,
		Logo:         form.Logo,
		CompType:     form.CompType,
		PeopleNum:    form.PeopleNum,
		City:         form.City,
		Introduction: form.Introduction,
		Address:      form.Address,
	}
	if form.Url != "" {
		comp.Url = form.Url
	}
	if err := db.Create(&comp).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return GetCompById(db, comp.Id)
}

func GetComp(db *gorm.DB, cond models.Comp) (*models.Comp, error) {
	comp := models.Comp{}
	if err := db.Where(cond).First(&comp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotExist
		}
		return nil, errors.AutoDbErr(err)
	}
	return &comp, nil
}

func GetCompById(db *gorm.DB, id models.Id) (*models.Comp, error) {
	cond := models.Comp{}
	cond.Id = id
	return GetComp(db, cond)
}

func GetCompByName(db *gorm.DB, name string) (*models.Comp, error) {
	return GetComp(db, models.Comp{CompName: name})
}

func QueryCompWithCond(db *gorm.DB, cond *models.Comp) *gorm.DB {
	return db.Model(&models.Comp{}).Where(cond)
}

func QueryComp(db *gorm.DB, q string) *gorm.DB {
	return db.Model(&models.Comp{}).Where("comp_name LIKE ?", "%"+q+"%").Order("created_at")
}

func CloseComp(db *gorm.DB, compId models.Id) error {
	if err := db.Model(&models.Comp{}).Where("id", compId).Update("state", "inactive").Error; err != nil {
		return errors.AutoDbErr(err)
	}
	return nil
}

func UpdateComp(db *gorm.DB, form *models.Comp) (*models.Comp, error) {
	if err := db.Model(&models.Comp{}).Where("id", form.Id).Updates(&form).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return nil, nil
}
