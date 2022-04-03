package web

import (
	"fmt"
	"net/http"
	"strings"

	_ "CampusRecruitment/docs"
	"CampusRecruitment/pkg/apps/ctx"
	"CampusRecruitment/pkg/config"
	"CampusRecruitment/pkg/consts"
	"CampusRecruitment/pkg/infra/log"
	"CampusRecruitment/pkg/web/apiv1"
	"github.com/gin-gonic/gin"
	jsoniterExtra "github.com/json-iterator/go/extra"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	// lowerCamelCase
	jsoniterExtra.SetNamingStrategy(func(name string) string {
		return fmt.Sprintf("%s%s", strings.ToLower(string(name[0])), name[1:])
	})
}

func Start() error {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.CustomRecoveryWithWriter(log.Writer(), func(c *gin.Context, r interface{}) {
		ac := ctx.New(c)
		if err, ok := r.(error); ok {
			ac.Response(nil, err)
		} else {
			ac.Response(nil, fmt.Errorf("%v", r))
		}
	}))

	e.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"version": consts.VERSION,
			"commit":  consts.COMMIT,
		})
	})
	// swagger api
	e.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// http api
	apiv1.Register(e.Group("/api/v1"))

	log.Get().Infof("Listening and serving HTTP on %s", config.Get().Listen)
	return e.Run(config.Get().Listen)
}
