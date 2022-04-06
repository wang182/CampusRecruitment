package models

import "CampusRecruitment/pkg/consts"

// UserId 同步自 iac
type User struct {
	TimedModel

	Email    string `gorm:"uniqueIndex:idx_user_email;not null;" json:"email"`
	Password string `gorm:"not null;" json:"-"`
	Name     string `gorm:"not null;" json:"name"`
	Phone    string `gorm:"not null;" json:"phone"`
	From     string `gorm:"not null;" json:"from"`
	Position string `gorm:"" json:"position"`
	Sex      string `gorm:"type:enum('man','woman');not null" json:"sex"`
	Role     string `gorm:"type:enum('admin','stu','comp');default:stu;not null" json:"role"`
	HeadImg  string `gorm:"not null" json:"headImg"`
}

// 如果需要可以定义 TableName() 函数返回自定义表名，否则会自动设置表名
func (User) TableName() string {
	return consts.TablePrefix + "user"
}
