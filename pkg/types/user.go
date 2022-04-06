package types

import (
	"CampusRecruitment/pkg/models"

	"github.com/dgrijalva/jwt-go"
)

type UserRegisterForm struct {
	baseForm

	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	HeadImg  string `json:"headImg" form:"headImg" binding:""`
	Name     string `json:"name" form:"name" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	From     string `json:"from" form:"from" binding:"required"`
	Position string `json:"position" form:"position" binding:""`
	Role     string `json:"role" form:"role" binding:"required"`
	Sex      string `json:"sex" form:"sex" binding:"required"`
}

type UserLoginForm struct {
	baseForm

	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserTokenClaims struct {
	jwt.StandardClaims

	UserId models.Id `json:"userId"`
	Email  string    `json:"email"`
}

type SearchUserForm struct {
	pageForm

	Q string `json:"q" form:"q"` // 查询关键字
}

type GetUserForm struct {
	baseForm

	Id models.Id `uri:"id"`
}

type UpdateUserPassword struct {
	baseForm
	UserId  models.Id `json:"id" form:"id" swaggerignore:"true"`
	OldPass string    `json:"oldPass" form:"oldPass" binding:"required"`
	NewPass string    `json:"newPass" form:"oldPass" binding:"required"`
}

type NormalUpdateUserForm struct {
	baseForm
	UserId   models.Id `json:"id" form:"id" swaggerignore:"true"`
	Name     string    `json:"name" form:"name"`
	Phone    string    `json:"phone" form:"phone"`
	From     string    `json:"from" form:"from"`
	Sex      string    `json:"sex" form:"sex"`
	Position string    `json:"position" form:"position"`
	HeadImg  string    `json:"headImg" form:"headImg"`
}

type UpdateUserForm struct {
	baseForm
	UserId   models.Id `json:"id" form:"id" swaggerignore:"true"`
	Name     string    `json:"name" form:"name"`
	Phone    string    `json:"phone" form:"phone"`
	From     string    `json:"from" form:"from"`
	Sex      string    `json:"sex" form:"sex"`
	HeadImg  string    `json:"headImg" form:"headImg"`
	Password string    `json:"password" form:"password"`
	Position string    `json:"position" form:"position"`
	Role     string    `json:"role" form:"role"`
}
