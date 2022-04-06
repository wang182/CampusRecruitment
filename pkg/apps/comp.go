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
	if err := query.Find(&comps).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &comps)
}

func QueryComp(c *ctx.Context, form *types.QueryCompForm) (interface{}, error) {
	query := services.QueryComp(c.DB(), form.Q)
	comps := make([]resps.CompCondResp, 0)
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
	//todo 修改招聘岗位状态
	//if err = services.UpdateJobState(c.DB(), comp.Id); err != nil {
	//	return nil, err
	//}
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
	form := types.UpdateCompForm{
		State: "active",
	}
	form.CompId = compId
	comp, err = services.UpdateComp(c.DB(), &form)
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
	comp, err := services.GetCompById(c.DB(), form.CompId)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	if comp == nil {
		return nil, errors.ErrNotExist
	}
	if user.From != comp.CompName {
		return nil, errors.ErrPermDeny
	}
	comp, err = services.UpdateComp(c.DB(), form)
	if err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return comp, nil
}
