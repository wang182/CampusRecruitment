package types

import (
	"CampusRecruitment/pkg/models"

	"github.com/dgrijalva/jwt-go"
)

type UserRegisterForm struct {
	baseForm

	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserLoginForm struct {
	baseForm

	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserTokenClaims struct {
	jwt.StandardClaims

	UserId   models.Id `json:"userId"`
	UserName string    `json:"userName"`
}

type UserLoginResp struct {
	Token string `json:"token"`
}

type SearchUserForm struct {
	pageForm

	Q string `json:"q" form:"q"` // 查询关键字
}

type GetUserForm struct {
	baseForm

	Id models.Id `uri:"id"`
}
