package apps

import (
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/services"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"CampusRecruitment/pkg/types/resps"
	"time"
)

func UserRegister(c *ctx.Context, form *types.UserRegisterForm) (interface{}, error) {
	if form.Role == "comp" {
		comp, err := services.GetCompByName(c.DB(), form.From)
		if err != nil {
			if err == errors.ErrNotExist {
				return nil, errors.WarnCompNotExit
			} else {
				return nil, err
			}
		}
		if comp.State == "approve" {
			return nil, errors.WarnCompNotActive
		}
	}
	user, err := services.CreateUser(c.DB(), form)
	if err != nil {
		return nil, err
	}
	c.Logger().Infof("created user: %s(%s)", user.Email, user.Id)
	return user, nil
}

func UserLogin(c *ctx.Context, form *types.UserLoginForm) (interface{}, error) {
	lg := c.Logger()

	user, err := services.GetUserByEmail(c.DB(), form.Email)
	if err != nil {
		if errors.IsNotFoundErr(err) {
			return nil, errors.ErrAuthFailed
		}
		return nil, err
	}

	// 验证
	ok, err := services.VerifyUserPassword(form.Password, user.Password)
	if err != nil {
		lg.Errorf("verify user password: %v", err)
		return nil, err
	}
	if !ok {
		return nil, errors.ErrAuthFailed
	}

	// 发行自己的 jwt token
	token, err := services.GenerateJwtToken(user.Id, user.Email, 24*time.Hour)
	if err != nil {
		lg.Errorf("SsoLogin -> GenerateJwtToken: %s", err.Error())
		return nil, err
	}

	return &resps.UserLoginResp{
		Token: token,
	}, nil
}

func SearchUser(c *ctx.Context, form *types.SearchUserForm) (interface{}, error) {
	query := services.QueryUser(c.DB(), form.Q)
	users := make([]models.User, 0)
	if err := query.Find(&users).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return c.PageResult(query, form.Page(), form.PageSize(), &users)
}

func AuthMe(c *ctx.Context) (interface{}, error) {
	userInfo, err := c.User()
	if err != nil {
		return nil, err
	}
	user, err := services.GetUserById(c.DB(), userInfo.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(c *ctx.Context, userId models.Id) (interface{}, error) {
	userinfo, err := c.User()
	if err != nil {
		return nil, err
	}
	//删之前确认一下id是否存在
	_, err = services.GetUserById(c.DB(), userId)
	if err != nil {
		return nil, err
	}
	//如果删除用户不是自己（注销），则判断是否为管理员权限
	if userinfo.Id != userId {
		//判断是否有删除权限
		if err := services.HasDeleteUserPerm(c.DB(), userinfo.Id); err != nil {
			return nil, err
		}
	}

	if err := services.DeleteUser(c.DB(), userId); err != nil {
		return nil, err
	}
	return nil, nil
}

func UpdateUserPassword(c *ctx.Context, form *types.UpdateUserPassword) (interface{}, error) {
	lg := c.Logger()
	user, err := services.GetUserById(c.DB(), form.UserId)
	if err != nil {
		return nil, err
	}
	//比较旧密码正确性
	ok, verifyErr := services.VerifyUserPassword(form.OldPass, user.Password)
	if verifyErr != nil {
		lg.Errorf("verify user password: %v", err)
		return nil, verifyErr
	}
	if !ok {
		return nil, errors.ErrAuthFailed
	}
	userInfo, updateErr := services.UpdatePass(c.DB(), form.NewPass)
	if updateErr != nil {
		return nil, updateErr
	}
	return resps.PassUpdateResp{
		Id:    userInfo.Id,
		Role:  userInfo.Role,
		Email: userInfo.Email,
		Name:  userInfo.Name,
	}, nil
}

func NormalUserUpdate(c *ctx.Context, form *types.NormalUpdateUserForm) (interface{}, error) {
	user, err := services.GetUserById(c.DB(), c.MustUser().Id)
	if err != nil {
		return nil, err
	}
	if user.Role == "comp" {
		if form.From != "" {
			comp, err := services.GetCompByName(c.DB(), form.From)
			if err != nil {
				return nil, err
			}
			if comp == nil {
				return nil, errors.WarnCompNotExit
			}
		}
	}
	user, err = services.NormalUserUpdate(c.DB(), form)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AdminUserUpdate(c *ctx.Context, form *types.UpdateUserForm) (interface{}, error) {
	user, err := services.GetUserById(c.DB(), c.MustUser().Id)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, errors.ErrAuthFailed
	}
	if form.From != "" {
		comp, err := services.GetCompByName(c.DB(), form.From)
		if err != nil {
			return nil, err
		}
		if comp == nil {
			return nil, errors.WarnCompNotExit
		}
	}
	userInfo, err := services.UpdateUser(c.DB(), form)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func CheckLoginUserRole(c *ctx.Context) (string, error) {
	user, err := services.GetUserById(c.DB(), c.MustUser().Id)
	if err != nil {
		return "", errors.AutoDbErr(err)
	}
	return user.Role, nil
}
