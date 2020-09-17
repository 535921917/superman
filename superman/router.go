package main

import (
	"github.com/gin-gonic/gin"
	"shuai/superman/manage"
)

func RegisterRouter(engine *gin.Engine) {

	engine.GET("/api/service/register", manage.RegisterOrKeep)
	engine.GET("/api/service/cancel", manage.Cancel)
	engine.GET("/api/service/list", manage.GetServiceList)
	engine.GET("/api/service/version", manage.GetServiceVersion)


}
