package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model/response"
)

type UserController struct {
	JWTRequired bool
}

func (uc *UserController) PostUserInfo(c *gin.Context) {
	response.Response(c, http.StatusOK, "", gin.H{
		"message": "获取用户信息",
	})
}
