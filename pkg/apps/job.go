package apps

import (
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/services"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"CampusRecruitment/pkg/types/resps"
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

func QueryJobs(c *ctx.Context, form *types.QueryJobsForm) (interface{}, error) {
	query := services.QueryJobs(c.DB(), form.Q)
	jobs := make([]models.Job, 0)
	if err := query.Where("state", "active").Find(&jobs).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &jobs)
}

func AdminQueryJobs(c *ctx.Context, form *types.QueryJobsForm) (interface{}, error) {
	query := services.QueryJobs(c.DB(), form.Q)
	jobs := make([]models.Job, 0)
	if err := query.Find(&jobs).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &jobs)
}

func QueryJobsWithCond(c *ctx.Context, form *types.SearchJobsCondForm) (interface{}, error) {
	job := types.JobCondForm{}
	comp := types.QueryCompWithCondForm{}
	if form.Wage != "" {
		job.WageSection = form.Wage
	}
	if form.City != "" {
		job.City = form.City
	}
	if form.CompType != "" {
		comp.CompType = form.CompType
	}
	if form.CompPeople != "" {
		comp.PeopleNum = form.CompPeople
	}
	ids, err := SearchCompIdsWithCond(c, &comp)
	if err != nil {
		return nil, err
	}
	query := services.SearchJobsWithCond(c.DB(), &job)
	if len(ids) > 0 {
		query.Where("comp_id In ?", ids)
	}
	jobs := make([]models.Job, 0)
	if err = query.Scan(&jobs).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), jobs)
}

func CloseJob(c *ctx.Context, id models.Id) (interface{}, error) {
	role, err := CheckLoginUserRole(c)
	if err != nil {
		return nil, err
	}
	if role == "stu" {
		return nil, errors.ErrPermDeny
	}
	job, err := services.GetJobById(c.DB(), id)
	if err != nil {
		return nil, err
	}
	if job == nil {
		return nil, errors.ErrNotExist
	}
	if err = services.CLoseJob(c.DB(), id); err != nil {
		return nil, err
	}
	return nil, nil
}

func DeleteJob(c *ctx.Context, id models.Id) (interface{}, error) {
	role, err := CheckLoginUserRole(c)
	if err != nil {
		return nil, err
	}
	if role != "admin" {
		return nil, errors.ErrPermDeny
	}
	job, err := services.GetJobById(c.DB(), id)
	if err != nil {
		return nil, err
	}
	if job == nil {
		return nil, errors.ErrNotExist
	}
	if err = services.DeleteJob(c.DB(), id); err != nil {
		return nil, err
	}
	return nil, nil
}

func CloseJobsByIds(c *ctx.Context, form *types.CloseJobsForm) (interface{}, error) {
	role, err := CheckLoginUserRole(c)
	if err != nil {
		return nil, err
	}
	if role == "stu" {
		return nil, errors.ErrPermDeny
	}

	if err = services.UpdateJobsState(c.DB(), form.Ids, "inactive"); err != nil {
		return nil, err
	}
	return nil, nil
}

func GetJobDetails(c *ctx.Context, form *types.GetJobDetailsForm) (interface{}, error) {
	job, err := services.GetJobById(c.DB(), form.JobId)
	if err != nil {
		return nil, err
	}
	jobResp := resps.JobRespDetail{
		JobName: job.JobName,
		JobNum:  job.JobNum,
		MinWage: job.MinWage,
		MaxWage: job.MaxWage,
		Desc:    job.Desc,
		City:    job.City,
		Address: job.Address,
		Tags:    services.ChangeStringToArray(job.Tags),
	}
	user, err := services.GetUserById(c.DB(), form.PublishId)
	if err != nil {
		return nil, err
	}
	userResp := resps.JobDetailsUserResp{
		Position: user.Position,
		HeadImg:  user.HeadImg,
		Name:     user.Name,
	}
	comp, err := services.GetCompById(c.DB(), form.CompId)
	if err != nil {
		return nil, err
	}
	compResp := resps.CompCondResp{
		CompName:  comp.CompName,
		CompType:  comp.CompType,
		Logo:      comp.Logo,
		PeopleNum: comp.PeopleNum,
	}
	resp := resps.GetJobDetailResp{
		JobRespDetail:      jobResp,
		JobDetailsUserResp: userResp,
		CompCondResp:       compResp,
	}
	return resp, nil
}

func ApproveJob(c *ctx.Context, jobId models.Id) (interface{}, error) {
	role, err := CheckLoginUserRole(c)
	if err != nil {
		return nil, err
	}
	if role != "admin" {
		return nil, errors.ErrPermDeny
	}
	ids := make([]models.Id, 0)
	ids[0] = jobId
	if err = services.UpdateJobsState(c.DB(), ids, "active"); err != nil {
		return nil, err
	}
	return nil, nil
}
