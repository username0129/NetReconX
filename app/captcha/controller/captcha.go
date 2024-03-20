package controller

import (
	"errors"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
	"server/app/captcha/model/response"
	"server/internal/e"
	"server/internal/global"
	"server/internal/model/common"
	"server/internal/utils"
)

type CaptchaController struct{}

func (cc *CaptchaController) PostCaptcha(c *gin.Context) {
	openCaptcha := global.Config.Captcha.OpenCaptcha               // 是否开启验证码
	openCaptchaTimeOut := global.Config.Captcha.OpenCaptchaTimeOut // 是否开启验证码

	key := c.ClientIP() // 客户端 IP

	item, err := utils.GetCacheItem(key)
	if err != nil {
		// 当条目不存在时或者超时时，初始化条目
		if errors.Is(err, bigcache.ErrEntryNotFound) || errors.Is(err, e.ErrCacheEntryTimeout) {
			utils.SetCacheItem(key, []byte("1"), openCaptchaTimeOut)
		} else {
			global.Logger.Error("获取缓存条目错误", zap.Error(err))
			return
		}
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < utils.ItemToInt(item) {
		oc = true
	}

	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.Long, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, global.CaptchaStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.Logger.Error("验证码获取失败", zap.Error(err))
		common.Response(c, http.StatusInternalServerError, "验证码获取失败", nil)
		return
	}

	common.Response(c, http.StatusOK, "验证码获取成功", response.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.Config.Captcha.Long,
		OpenCaptcha:   oc,
	})
}
