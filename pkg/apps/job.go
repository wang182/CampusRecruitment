package apps

import (
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/services"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
)

func CreateJob(c *ctx.Context, form *types.CreateJobForm) (interface{}, error) {
	user, err := services.GetUserById(c.DB(), c.MustUser().Id)
	if err != nil {
		return nil, err
	}
	if user.Role != "comp" {
		return nil, errors.ErrPermDeny
	}
	form.PublishId = user.Id
	comp, err := services.GetCompByName(c.DB(), user.From)
	if err != nil {
		return nil, err
	}
	form.CompId = comp.Id
	job, err := services.CreateJob(c.DB(), form)
	if err != nil {
		return nil, err
	}
	return job, nil
}
