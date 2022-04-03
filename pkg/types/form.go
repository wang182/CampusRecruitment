package types

import (
	"CampusRecruitment/pkg/consts"
)

type Form interface {
	Validate() error
}

type baseForm struct{}

func (form *baseForm) Validate() error {
	return nil
}

type pageForm struct {
	baseForm

	Page_     int `json:"page" form:"page"`
	PageSize_ int `json:"pageSize" form:"pageSize"`
}

func (f *pageForm) Page() int {
	if f.Page_ < 1 {
		return 1
	}
	return f.Page_
}

func (f *pageForm) PageSize() int {
	if f.PageSize_ > consts.MaxPageSize {
		return consts.MaxPageSize
	} else if f.PageSize_ <= 0 {
		return consts.DefaultPageSize
	}
	return f.PageSize_
}
