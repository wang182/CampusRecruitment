package handler

import (
	"CampusRecruitment/pkg/apps"
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/types"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @Tags Users
// @Summary 用户注册
// @Accept json
// @Produce json
// @Param  form body types.UserRegisterForm true "parameter"
// @router /users/register [POST]
// @Success 200 {object} types.APIResponse{result=models.User}
func UserRegister(c *gin.Context) {
	ac := ctx.New(c)
	form := types.UserRegisterForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}

	ac.Response(apps.UserRegister(ac, &form))
}

// UserLogin 用户登录
// @Tags Users
// @Summary 用户登录
// @Accept json
// @Produce json
// @Param json body types.UserLoginForm true "parameter"
// @router /users/login [POST]
// @Success 200 {object} types.APIResponse{result=resps.UserLoginResp}
func UserLogin(c *gin.Context) {
	ac := ctx.New(c)
	form := types.UserLoginForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}

	ac.Response(apps.UserLogin(ac, &form))
}

// SearchUser 查询用户
// @Tags Users
// @Summary 查询用户列表
// @Accept application/x-www-form-urlencoded, application/json
// @Produce json
// @Security AuthToken
// @Param json query types.SearchUserForm true "parameter"
// @router /users [GET]
// @Success 200 {object} types.APIResponse{result=types.APIPageResult}
func SearchUser(c *gin.Context) {
	ac := ctx.New(c)
	form := types.SearchUserForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}

	ac.Response(apps.SearchUser(ac, &form))
}

// AuthMe 获取登录信息
// @Tags Users
// @Summary 获取登录信息
// @Accept json
// @Produce json
// @Security AuthToken
// @router /users/auth [GET]
// @Success 200 {object} types.APIResponse{result=models.User}
func AuthMe(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(apps.AuthMe(ac))
}

// UserDelete 用户删除
// @Tags Users
// @Summary 用户删除
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path string true "user id"
// @router /users/{id} [DELETE]
// @Success 200 {object} types.APIResponse
func UserDelete(c *gin.Context) {
	ac := ctx.New(c)
	ac.Response(apps.DeleteUser(ac, models.Id(c.Param("id"))))
}

// UpdatePassword 普通用户修改密码
// @Tags Users
// @Summary 普通用户修改密码
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json body types.UpdateUserPassword true "parameter"
// @router /users/updatePass [PUT]
// @Success 200 {object} types.APIResponse{result=resps.PassUpdateResp}
func UpdatePassword(c *gin.Context) {
	ac := ctx.New(c)
	form := types.UpdateUserPassword{}
	form.UserId = ac.MustUser().Id
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.UpdateUserPassword(ac, &form))
}

// NormalUserUpdate 用户修改
// @Tags Users
// @Summary 用户修改
// @Accept json
// @Produce json
// @Security AuthToken
// @Param json body types.NormalUpdateUserForm true "parameter"
// @router /users [PUT]
// @Success 200 {object} types.APIResponse{result=models.User}
func NormalUserUpdate(c *gin.Context) {
	ac := ctx.New(c)
	form := types.NormalUpdateUserForm{}
	form.UserId = ac.MustUser().Id
	if err := ac.Bind(&form); err != nil {
		return
	}

	ac.Response(apps.NormalUserUpdate(ac, &form))
}

// UserUpdate 管理员修改用户
// @Tags Users
// @Summary 管理员修改用户
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path string true "user id"
// @Param json body types.UpdateUserForm true "parameter"
// @router /users/adminUpdate/{id} [PUT]
// @Success 200 {object} types.APIResponse{result=models.User}
func UserUpdate(c *gin.Context) {
	ac := ctx.New(c)
	form := types.UpdateUserForm{}
	form.UserId = models.Id(c.Param("id"))
	if err := ac.Bind(&form); err != nil {
		return
	}
	ac.Response(apps.AdminUserUpdate(ac, &form))
}
