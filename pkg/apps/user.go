package apps

import (
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/services"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"time"
)

func UserRegister(c *ctx.Context, form *types.UserRegisterForm) (interface{}, error) {
	user, err := services.CreateUser(c.DB(), form.Username, form.Password)
	if err != nil {
		return nil, err
	}
	c.Logger().Infof("created user: %s(%s)", user.Username, user.Id)
	return nil, nil
}

func UserLogin(c *ctx.Context, form *types.UserLoginForm) (interface{}, error) {
	lg := c.Logger()

	user, err := services.GetUserByName(c.DB(), form.Username)
	if err != nil {
		if errors.IsNotFoundErr(err) {
			return nil, errors.ErrAuthFailed
		}
		return nil, err
	}

	// 验证token by cloudiac api
	ok, err := services.VerifyUserPassword(form.Password, user.Password)
	if err != nil {
		lg.Errorf("verify user password: %v", err)
		return nil, err
	}
	if !ok {
		return nil, errors.ErrAuthFailed
	}

	// 发行自己的 jwt token
	token, err := services.GenerateJwtToken(user.Id, user.Username, 24*time.Hour)
	if err != nil {
		lg.Errorf("SsoLogin -> GenerateJwtToken: %s", err.Error())
		return nil, err
	}

	return &types.UserLoginResp{
		Token: token,
	}, nil
}

func SearchUser(c *ctx.Context, form *types.SearchUserForm) (interface{}, error) {
	c.User()

	query := services.QueryUser(c.DB(), form.Q)
	users := make([]models.User, 0)
	if err := query.Find(&users).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &users)
}
