package apiv1

import (
	"CampusRecruitment/pkg/web/apiv1/handler"
	"CampusRecruitment/pkg/web/middleware"

	"github.com/gin-gonic/gin"
)

// @title 校友招聘系统
// @version 1.0.0
// @description 校友招聘系统接口文档

// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization

func Register(g *gin.RouterGroup) {
	// 用户注册
	g.POST("/users/register", handler.UserRegister)
	// 用户登录
	g.POST("/users/login", handler.UserLogin)

	{
		//文件上传
		g.POST("icons", handler.IconFileUpload)
		g.GET("/icons", handler.ReadIconFile)
	}
	// add before apis which need jwt auth
	g.Use(middleware.JwtAuth())
	g.GET("/users", handler.SearchUser)
	// g.GET("/users/:id", handler.GetUser)
}
