package apps

import (
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/services"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"CampusRecruitment/pkg/types/resps"
)

func CompRegister(c *ctx.Context, form *types.CompRegisterForm) (interface{}, error) {
	comp, err := services.CreateComp(c.DB(), form)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func QueryCompWithCond(c *ctx.Context, form *types.QueryCompWithCondForm) (interface{}, error) {
	cond := models.Comp{}
	if form.CompType != "" {
		cond.CompType = form.CompType
	}
	if form.City != "" {
		cond.City = form.City
	}
	if form.PeopleNum != "" {
		cond.PeopleNum = form.PeopleNum
	}
	query := services.QueryCompWithCond(c.DB(), &cond)
	comps := make([]resps.CompCondResp, 0)
	if err := query.Where("state", "active").Find(&comps).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &comps)
}

func SearchCompIdsWithCond(c *ctx.Context, form *types.QueryCompWithCondForm) ([]string, error) {
	cond := models.Comp{}
	if form.CompType != "" {
		cond.CompType = form.CompType
	}
	if form.City != "" {
		cond.City = form.City
	}
	if form.PeopleNum != "" {
		cond.PeopleNum = form.PeopleNum
	}
	ids := make([]string, 0)
	if err := services.QueryCompWithCond(c.DB(), &cond).Select("id").Find(&ids).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return ids, nil
}

func QueryComp(c *ctx.Context, form *types.QueryCompForm) (interface{}, error) {
	query := services.QueryComp(c.DB(), form.Q)
	comps := make([]resps.CompCondResp, 0)
	if err := query.Where("state", "active").Find(&comps).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &comps)
}

func AdminQueryComp(c *ctx.Context, form *types.QueryCompForm) (interface{}, error) {
	query := services.QueryComp(c.DB(), form.Q)
	comps := make([]models.Comp, 0)
	if err := query.Find(&comps).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &comps)
}

func CloseComp(c *ctx.Context, id models.Id) (interface{}, error) {
	user, err := services.GetUserById(c.DB(), c.MustUser().Id)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	if user.Role != "admin" {
		return nil, errors.ErrPermDeny
	}
	comp, err := services.GetCompById(c.DB(), id)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	if comp == nil {
		return nil, errors.ErrNotExist
	}
	//更改所有公司员工在职状态
	if err = services.UpdateUserState(c.DB(), comp.CompName); err != nil {
		return nil, err
	}
	//关闭公司更改职位状态为inactive
	jobsId, err := services.GetJobsIdByCompId(c.DB(), comp.Id)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	if err = services.UpdateJobsState(c.DB(), jobsId, "inactive"); err != nil {
		return nil, err
	}
	if err = services.CloseComp(c.DB(), comp.Id); err != nil {
		return nil, err
	}
	return nil, nil
}

func ApproveComp(c *ctx.Context, compId models.Id) (interface{}, error) {
	user, err := services.GetUserById(c.DB(), c.MustUser().Id)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	if user.Role != "admin" {
		return nil, errors.ErrPermDeny
	}
	comp, err := services.GetCompById(c.DB(), compId)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	if comp == nil {
		return nil, errors.ErrNotExist
	}
	compForm := models.Comp{
		State: "active",
	}
	compForm.Id = compId
	comp, err = services.UpdateComp(c.DB(), &compForm)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func UpdateComp(c *ctx.Context, form *types.UpdateCompForm) (interface{}, error) {
	user, err := services.GetUserById(c.DB(), c.MustUser().Id)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	comp, err := services.GetCompById(c.DB(), form.Id)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	if comp == nil {
		return nil, errors.ErrNotExist
	}
	if user.From != comp.CompName {
		return nil, errors.ErrPermDeny
	}

	comp, err = services.UpdateComp(c.DB(), ChangeUpdateCompFormToComp(form))
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return comp, nil
}

func ChangeUpdateCompFormToComp(form *types.UpdateCompForm) *models.Comp {
	compForm := models.Comp{}
	compForm.Id = form.Id
	if form.City != "" {
		compForm.City = form.City
	}
	if form.CompType != "" {
		compForm.CompType = form.CompType
	}
	if form.CompName != "" {
		compForm.CompName = form.CompName
	}
	if form.State != "" {
		compForm.State = form.State
	}
	if form.PeopleNum != "" {
		compForm.PeopleNum = form.PeopleNum
	}
	if form.Address != "" {
		compForm.Address = form.Address
	}
	if form.Url != "" {
		compForm.Url = form.Url
	}
	if form.Introduction != "" {
		compForm.Introduction = form.Introduction
	}
	if form.Logo != "" {
		compForm.Logo = form.Logo
	}
	return &compForm
}

func GetCompDetail(c *ctx.Context, compId models.Id) (interface{}, error) {
	hotJobs, err := services.GetHotJobsByCompId(c.DB(), compId)
	if err != nil {
		return nil, err
	}
	comp, err := services.GetCompById(c.DB(), compId)
	if err != nil {
		return nil, err
	}
	userNum, err := services.GetUserNumByCompId(c.DB(), compId)
	if err != nil {
		return nil, err
	}
	CompDetailResp := resps.CompDetailResp{
		CompName:     comp.CompName,
		Logo:         comp.Logo,
		CompType:     comp.CompType,
		PeopleNum:    comp.PeopleNum,
		JobNum:       services.GetJobNumByCompId(c.DB(), compId),
		UserNum:      userNum,
		City:         comp.City,
		Introduction: comp.Introduction,
		Address:      comp.Address,
		Url:          comp.Url,
		HotJobs:      hotJobs,
	}
	return CompDetailResp, nil
}
