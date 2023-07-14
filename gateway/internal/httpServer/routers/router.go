package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/adminToServer", admin.Test)
	router.POST("/adminToServer", admin.MsgHandler)
	return router
}
