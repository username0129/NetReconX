package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

// PostLogin
//
//	@Description: 完成用户登录
//	@receiver ac
//	@Router    /auth/login
func (ac *AuthController) PostLogin(c *gin.Context) {
	// 示例：
	c.JSON(http.StatusOK, gin.H{"message": "完成用户登录"})
}
