package core

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery()) // 避免 panic 导致服务器停止

	return router
}
