package models

type Job struct {
	TimedModel
	JobName     string   `gorm:"jobName;not null;"`
	PublishId   Id       `gorm:"not null"`
	CompId      Id       `gorm:"not null"`
	MinWage     int      `gorm:"type:enum('unlimited','3k-','3-5k','5-10k','10-15k','15-20k','20k+');not null"`
	MaxWage     int      `gorm:"maxWage"`
	WageSection string   `json:"wageSection"`
	JobNum      int      `gorm:"not null"`
	Desc        string   `gorm:"type:text;not null;"`
	Address     string   `gorm:"not null"`
	Tags        []string `gorm:""`
	State       string   `gorm:"type:enum('approving','active','inactive');default:approving"`
}
