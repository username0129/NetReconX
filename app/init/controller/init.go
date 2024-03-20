package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/app/init/model"
	"server/app/init/service"
	"server/internal/global"
	"server/internal/model/common"
)

type InitController struct{}

var initService = new(service.InitService)

// PostInit
//
//	@Description: 初始化数据库
//	@receiver ic
//	@param c
//	@Router: /init/init
func (ic *InitController) PostInit(c *gin.Context) {
	if global.DB != nil {
		global.Logger.Error("已存在数据库配置")
		common.Response(c, http.StatusInternalServerError, "已存在数据库配置", nil)
		return
	}

	var req model.InitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Error("参数解析错误", zap.Error(err))
		common.Response(c, http.StatusInternalServerError, "参数解析错误", nil)
		return
	}

	if err := initService.Init(req); err != nil {
		global.Logger.Error("数据库初始化错误", zap.Error(err))
		common.Response(c, http.StatusInternalServerError, "数据库初始化错误", nil)
		return
	}
	common.Response(c, http.StatusOK, "数据库初始化成功", nil)
}
