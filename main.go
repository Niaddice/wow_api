package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/illidan33/wow_api/global"
	"github.com/illidan33/wow_api/modules"
	"github.com/illidan33/wow_api/routers"
	"github.com/illidan33/wow_api/routers/index"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	WowApi *gin.Engine
	DB     *gorm.DB
)

func main() {
	defer func() {
		modules.DbConn.Close()
	}()
	if global.Config.LogLevel != logrus.DebugLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	WowApi = gin.New()
	routers.Chart = WowApi.Group("/chart")
	routers.Chart.Use(index.AuthMiddleware)
	routers.Auth = WowApi.Group("/auth")
	routers.Auth.Use(index.AuthMiddleware)

	routers.Api = WowApi.Group("/api")
	routers.Macro = WowApi.Group("/macro")
	routers.MacroOld60 = WowApi.Group("/macro60")
	routers.New()
	WowApi.StaticFS("/js/", http.Dir("./public/js"))
	WowApi.StaticFS("/css/", http.Dir("./public/css"))
	WowApi.StaticFile("/favicon.ico", "./public/favicon.ico")
	WowApi.LoadHTMLGlob("./public/html/*/*")

	//WowApi.GET("/", index.Index)
	// 1. 将根路径 "/" 重定向到目标 URL
	WowApi.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/macro/view/combineSkills")
	})
	WowApi.NoRoute(index.Index)
	WowApi.Run(fmt.Sprintf("%s:%d", global.Config.ListenHost, global.Config.ListenPort))
}
