package models

type Comp struct {
	TimedModel
	CompName     string `gorm:"uniqueIndex;not null;"`
	Logo         string `gorm:"not null;"`
	CompType     string `gorm:"type:enum('mall','game','medical','hardware','software','network','finance','video','education');not null;"`
	PeopleNum    string `gorm:"type:enum('20','99','500','1000','9999','10000');"`
	City         string `gorm:"not null;"`
	Introduction string `gorm:"type:text;not null;"`
	Address      string `gorm:"not null;"`
	Url          string `gorm:"not null"`
	State        string `gorm:"type:enum('approve','active','inactive');default:approve;"`
}
