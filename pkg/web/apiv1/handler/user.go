package handler

import (
	"CampusRecruitment/pkg/apps"
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/types"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	ac := ctx.New(c)
	form := types.UserRegisterForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}

	ac.Response(apps.UserRegister(ac, &form))
}

func UserLogin(c *gin.Context) {
	ac := ctx.New(c)
	form := types.UserLoginForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}

	ac.Response(apps.UserLogin(ac, &form))
}

func SearchUser(c *gin.Context) {
	ac := ctx.New(c)
	form := types.SearchUserForm{}
	if err := ac.Bind(&form); err != nil {
		return
	}

	ac.Response(apps.SearchUser(ac, &form))
}
