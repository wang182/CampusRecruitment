package models

// UserId 同步自 iac
type User struct {
	TimedModel

	Username string `gorm:"uniqueIndex:idx_user_email;not null;"`
	Password string `gorm:"not null;" json:"-"`
}

// 如果需要可以定义 TableName() 函数返回自定义表名，否则会自动设置表名
// func (User) TableName() string {
// 	return consts.TablePrefix + "user"
// }
