package models

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"

	"CampusRecruitment/pkg/utils"
)

type BaseModel struct {
	Id Id `gorm:"primaryKey" json:"id"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if b.Id == "" {
		b.Id = Id(utils.GenId())
	}
	return nil
}

type UintIdModel struct {
	Id uint `gorm:"primarykey"`
}

type TimedModel struct {
	BaseModel

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type SoftDeleteModel struct {
	TimedModel

	DeletedAt soft_delete.DeletedAt `gorm:"index" json:"-"`
}

func Init(tx *gorm.DB) error {
	if err := migrate(tx); err != nil {
		return errors.Wrap(err, "models migrate")
	}
	return nil
}
