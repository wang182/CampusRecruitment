package apiv1

import (
	"CampusRecruitment/pkg/web/apiv1/handler"
	"CampusRecruitment/pkg/web/middleware"

	"github.com/gin-gonic/gin"
)

// @title 校园招聘系统
// @version 1.0.0
// @description 校园招聘系统接口文档

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
	{
		//用户相关api
		g.GET("/users/auth", handler.AuthMe)
		g.PUT("/users", handler.NormalUserUpdate)
		g.DELETE("/users/:id", handler.UserDelete)
		g.PUT("/users/updatePass", handler.UpdatePassword)
		g.GET("/users", handler.SearchUser)
		g.PUT("/users/adminUpdate/:id", handler.UserUpdate)

		g.POST("/comp", handler.CompRegister)
		g.GET("/comp/cond", handler.QueryCompWithCond)
		g.GET("/comp", handler.QueryComp)
		g.PUT("/comp/:id/close", handler.CloseComp)
		g.PUT("comp/:id/approve", handler.ApproveComp)
		g.PUT("/comp/:id", handler.UpdateComp)

		g.POST("/job", handler.CreateJob)

	}

}
