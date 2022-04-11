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
	//公司注册
	g.POST("/comp", handler.CompRegister)

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

		g.GET("/comp/cond", handler.QueryCompWithCond)
		g.GET("/comp", handler.QueryComp)
		g.GET("/admin/comp", handler.AdminQueryComp)
		g.PUT("/comp/:id/close", handler.CloseComp)
		g.PUT("comp/:id/approve", handler.ApproveComp)
		g.PUT("/comp/:id", handler.UpdateComp)
		g.GET("/comp/:id", handler.GetCompDetail)

		g.POST("/job", handler.CreateJob)
		g.GET("/jobs", handler.QueryJobs)
		g.GET("/admin/jobs", handler.AdminQueryJobs)
		g.GET("jobs/cond", handler.SearchJobsWithCond)
		g.PUT("/job/:id/close", handler.CloseJob)
		g.PUT("/job/ids/close", handler.CloseJobsByIds)
		g.DELETE("/job/:id", handler.DeleteJob)
		g.GET("/job/:id/details", handler.GetJobDetails)
		g.PUT("/job/:id", handler.ApproveJob)
	}

}
